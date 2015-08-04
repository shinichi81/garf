package cmd

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GarfCmd = &cobra.Command{
	Use:   "garf",
	Short: "GAPI is a simple API generator",
	Long:  `GAPI generates a simple API, structure and its modules, using Mux and Martini.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var Verbose, Overwrite bool
var Cwd, CfgFile, Seed, engine string
var garf *cobra.Command

func init() {
	GarfCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	GarfCmd.PersistentFlags().BoolVarP(&Overwrite, "force", "f", false, "force overwrite?")
	GarfCmd.PersistentFlags().StringVarP(&Cwd, "target", "t", "", "set target directory")
	GarfCmd.PersistentFlags().StringVarP(&Seed, "seed", "s", "", "set seed template directory")
	GarfCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "load build configuration from json file")
	GarfCmd.PersistentFlags().StringVarP(&engine, "engine", "e", "echo", "server engine (default: echo)")
	engine = strings.ToLower(engine)
	garf = GarfCmd
}

func InitializeConfig() {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		panic("Missing GOPATH on this environment.")
	}
	gopath, _ = filepath.EvalSymlinks(filepath.Join(gopath, "src/"))
	gopath, _ = filepath.Abs(gopath)

	viper.AddConfigPath(Cwd)

	if len(CfgFile) > 0 {
		viper.SetConfigFile(CfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	viper.Set("Bundles", make([]Bundle, 0))
	viper.SetDefault("Verbose", false)

	if garf.PersistentFlags().Lookup("verbose").Changed {
		viper.Set("Verbose", Verbose)
	}

	dir, _ := os.Getwd()
	if Cwd != "" {
		viper.Set("WorkingDir", filepath.Join(dir, Cwd))
	} else {
		viper.Set("WorkingDir", dir)
	}

	if Seed != "" {
		viper.Set("SeedDir", Seed)
	} else {
		_, filename, _, _ := runtime.Caller(1)
		dir := filepath.Join(path.Dir(filename), "../template")
		viper.Set("SeedDir", dir)
	}

	dir, _ = filepath.EvalSymlinks(viper.Get("WorkingDir").(string))
	dir, _ = filepath.Abs(dir)
	rootpath, _ := filepath.Rel(gopath, dir)
	viper.Set("rootpath", rootpath)
	viper.Set("engine", engine)
}

func Execute() {
	AddCommands()
	GarfCmd.Execute()
}

func AddCommands() {
	garf.AddCommand(newCmd)
}
