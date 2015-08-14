package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new...",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	newCmd.AddCommand(newApiCmd)
	newCmd.AddCommand(newBundleCmd)
}

var newApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Create new API structure and registry",
	Run:   NewApi,
}

func NewApi(cmd *cobra.Command, args []string) {
	InitializeConfig()

	cwd := viper.GetString("WorkingDir")
	path := cwd
	log.Println("Creating garf API in:", path)

	mkdir(cwd)
	mkdir(cwd, "garf")
	mkdir(cwd, "bundles")

	generate(".", "main")
	generate("garf", "garf")
}

var newBundleCmd = &cobra.Command{
	Use:   "bundle [name]",
	Short: "Create new API bundle.",
	Run:   NewBundle,
}

type Bundle struct {
	Name string
	Path string
}

func NewBundle(cmd *cobra.Command, args []string) {
	InitializeConfig()

	if len(args) < 1 {
		cmd.Usage()
		log.Fatalln("missing bundle name")
	}

	cwd := viper.GetString("WorkingDir")
	name := args[0]
	path := filepath.Join(cwd, "bundles", name)

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		cmd.Usage()
		log.Fatalln("Already exists", path)
	}

	addBundle(name)
}

func mkdir(x ...string) {
	p := filepath.Join(x...)

	log.Printf("Creating directory: %s\n", p)

	os.MkdirAll(p, 0777) // before umask
}

func generate(path, name string) {
	seedDir := viper.GetString("SeedDir")
	target := viper.GetString("WorkingDir")

	path = filepath.Join(target, path)

	tfile := filepath.Join(seedDir, fmt.Sprintf("%s%s", name, ".tmpl"))

	if _, err := os.Stat(tfile); os.IsNotExist(err) {
		log.Fatalln("Given seed directory does not contain template:", name)
		return
	}

	data, err := ioutil.ReadFile(tfile)
	if err != nil {
		log.Fatalln(tfile, "could not be loaded.")
		return
	}

	tmpl, _ := template.New(name).Parse(string(data))

	ffile := filepath.Join(path, fmt.Sprintf("%s%s", name, ".go"))
	var file *os.File
	file, err = os.Create(ffile)

	if err != nil {
		log.Fatal(ffile, "could not be created.")
		return
	}

	log.Printf("Generating file: %s\n", ffile)

	tmpl.Execute(file, viper.AllSettings())
}

func generateBundle(name string) {
	seedDir := viper.GetString("SeedDir")
	target := viper.GetString("WorkingDir")

	tdir := filepath.Join(seedDir, "bundle")

	if _, err := os.Stat(tdir); os.IsNotExist(err) {
		log.Fatalln("Given seed directory does not contain templates for bundles")
		return
	}

	bdir := filepath.Join(target, "bundles")

	mkdir(bdir, name)

	bdir = filepath.Join(bdir, name)

	nameRune := []rune(name)
	nameRune[0] = unicode.ToUpper(nameRune[0])
	capitalName := string(nameRune)
	lowerName := strings.ToLower(name)
	ctx := struct{ CapitalName, LowerName, engine string }{
		capitalName,
		lowerName,
		engine,
	}

	filepath.Walk(tdir, func(path string, f os.FileInfo, err error) error {
		tname := filepath.Base(path)
		tname = strings.Split(tname, ".")[0]

		if filepath.Ext(path) != ".tmpl" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		tmpl, err := template.New(fmt.Sprintf("%s%s", "template-", tname)).Parse(string(data))
		bfile := filepath.Join(bdir, fmt.Sprintf("%s%s", tname, ".go"))

		if err != nil {
			log.Fatalln(err)
			return err
		}

		log.Printf("Creating file: %s\n", bfile)

		var file *os.File
		file, err = os.Create(bfile)
		if err != nil {
			return nil
		}

		if err := tmpl.Execute(file, ctx); err != nil {
			log.Fatalln(err)
			return err
		}

		return nil
	})
}

func loadBundles() {
	target := viper.GetString("WorkingDir")

	tdir := filepath.Join(target, "bundles")

	if _, err := os.Stat(tdir); os.IsNotExist(err) {
		return
	}

	bund := viper.Get("Bundles").([]Bundle)
	filepath.Walk(tdir, func(path string, f os.FileInfo, err error) error {
		name := filepath.Base(path)

		if name == "bundles" {
			return nil
		}

		rpath, _ := filepath.Rel(target, path)
		if s, _ := os.Stat(path); s.IsDir() {
			bund = append(bund, Bundle{name, rpath})
		}
		return nil
	})

	viper.Set("Bundles", bund)
}

func addBundle(name string) {
	generateBundle(name)
	loadBundles()
	generate("garf", "garf")
}
