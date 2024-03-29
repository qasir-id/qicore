package scaffold

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

const (
	GoScaffoldPath = "src/github.com/qasir-id/qicore"
)

func init() {
	Gopath = os.Getenv("GOPATH")
	if Gopath == "" {
		panic("cannot find $GOPATH environment variable")
	}
}

var (
	Gopath        string
	ProjectName   string
	SubStrService string
)

type scaffold struct {
	debug bool
}

type data struct {
	AbsGenProjectPath string // The Abs Gen Project Path
	ProjectPath       string // The Go import project path (eg:github.com/fooOrg/foo)
	ProjectName       string // The project name which want to generated
	Quit              string
}

type templateEngine struct {
	Templates []templateSet
	currDir   string
}

type templateSet struct {
	templateFilePath string
	templateFileName string
	genFilePath      string
}

type DataFlag struct {
	Path          string
	Name          string
	SubStrService string
}

func New(debug bool) *scaffold {
	return &scaffold{debug: debug}
}

func (s *scaffold) Generate(dataFlag DataFlag) error {
	genAbsDir, err := filepath.Abs(dataFlag.Path)
	if err != nil {
		return err
	}
	// TODO: have to check path MUST be under the $GOPATH/src folder
	goProjectPath := strings.TrimPrefix(genAbsDir, filepath.Join(Gopath, "src")+string(os.PathSeparator))
	ProjectName = dataFlag.Name
	SubStrService = dataFlag.SubStrService
	d := data{
		AbsGenProjectPath: genAbsDir,
		ProjectPath:       goProjectPath,
		ProjectName:       ProjectName,
		Quit:              "<-quit",
	}

	if err := s.genFromTemplate(getTemplateSets(), d); err != nil {
		log.Println("error genFromTemplate", err)
		return err
	}

	return nil
}

func getTemplateSets() []templateSet {
	tt := templateEngine{}

	templateService := "/template/" + SubStrService + "/"
	templatesFolder := filepath.Join(Gopath, GoScaffoldPath, templateService)
	if os.Getenv("APP_MODE") == "develop" {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		templatesFolder = path + templateService
	}

	if err := filepath.Walk(templatesFolder, tt.visit); err != nil {
		log.Println("error get Template", err)
	}

	return tt.Templates
}

func (s *scaffold) genFromTemplate(templateSets []templateSet, d data) error {
	for _, tmpl := range templateSets {
		if err := s.tmplExec(tmpl, d); err != nil {
			return err
		}
	}
	return nil
}

func unescaped(x string) interface{} { return template.HTML(x) }

func (s *scaffold) tmplExec(tmplSet templateSet, d data) error {
	tmpl := template.New(tmplSet.templateFileName)
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": unescaped})
	tmpl, err := tmpl.ParseFiles(tmplSet.templateFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	relateDir := filepath.Dir(tmplSet.genFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(tmplSet.genFilePath))
	distAbsFilePath := filepath.Join(d.AbsGenProjectPath, distRelFilePath)
	s.debugPrintf("distRelFilePath:%s\n", distRelFilePath)
	s.debugPrintf("distAbsFilePath:%s\n", distAbsFilePath)
	if err := os.MkdirAll(filepath.Dir(distAbsFilePath), os.ModePerm); err != nil {
		return pkgErr.WithStack(err)
	}

	dist, err := os.Create(distAbsFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}

	defer dist.Close()

	fmt.Printf("Create %s\n", distRelFilePath)
	return tmpl.Execute(dist, d)
}

func (templEngine *templateEngine) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	projectName := ProjectName + "/"
	if ext := filepath.Ext(path); ext == ".tmpl" { // for handle file format tmpl
		templateFileName := filepath.Base(path)
		genFileBaeName := strings.TrimSuffix(templateFileName, ".tmpl") + ".go"
		genFileBasePath := filepath.Join(filepath.Dir(path), genFileBaeName)

		subStr := strings.Index(filepath.Dir(path), SubStrService)
		dirTemp := ""
		if subStr > -1 {
			dirTemp = genFileBasePath[subStr+7:]
		}

		projectName += dirTemp
		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      projectName,
		}
		templEngine.Templates = append(templEngine.Templates, templ)

	} else if mode := f.Mode(); mode.IsRegular() { // for handle file format except tmpl
		templateFileName := filepath.Base(path)
		basePath := filepath.Join(Gopath, GoScaffoldPath, "template")
		targPath := filepath.Join(filepath.Dir(path), templateFileName)
		genFileBasePath, err := filepath.Rel(basePath, targPath)
		if err != nil {
			return pkgErr.WithStack(err)
		}
		subStr := strings.Index(filepath.Dir(path), SubStrService+"/")
		dirTemp := ""
		if subStr > -1 {
			dirTemp = genFileBasePath[7:]
			projectName += dirTemp
		} else {
			projectName += dirTemp + "/" + templateFileName
		}
		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      projectName,
		}

		templEngine.Templates = append(templEngine.Templates, templ)
	}

	return nil
}

func (s *scaffold) debugPrintf(format string, a ...interface{}) {
	if s.debug == true {
		fmt.Printf(format, a...)
	}
}
