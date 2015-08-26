package main

import "fmt"
import "os"
import "strings"
import "io/ioutil"
import "path/filepath"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Localize(input string, dir string, output string, localization map[string]string) {
    files, _ := ioutil.ReadDir(input)
    for _, f := range files {
        if localization == nil {
            if f.IsDir() {
                Localize(fmt.Sprint(input + "/" + f.Name()), dir, output, localization)
            } else {
                data, error := ioutil.ReadFile(fmt.Sprint(input + "/" + f.Name()))
                if (error != nil) {
                    panic(error)
                }
                outputString := string(data)

                outputString = strings.Replace(outputString, "{{", "", -1)
                outputString = strings.Replace(outputString, "}}", "", -1)

                os.MkdirAll(strings.Replace(input, dir, output, 1), 0777)

                outputFile := strings.Replace(fmt.Sprint(input + "/" + f.Name()), dir, output, 1)

                f, err := os.Create(outputFile)
                check(err)
                defer f.Close()

                _, err = f.WriteString(outputString)
                f.Sync()
            }
        } else {
            if f.IsDir() {
                Localize(fmt.Sprint(input + "/" + f.Name()), dir, output, localization)
            } else {
                data, error := ioutil.ReadFile(fmt.Sprint(input + "/" + f.Name()))
                if (error != nil) {
                    panic(error)
                }
                outputString := string(data)

                for key := range localization {
                    outputString = strings.Replace(outputString, key, localization[key], -1)
                }

                os.MkdirAll(strings.Replace(input, dir, output, 1), 0777)

                outputFile := strings.Replace(fmt.Sprint(input + "/" + f.Name()), dir, output, 1)

                f, err := os.Create(outputFile)
                check(err)
                defer f.Close()

                _, err = f.WriteString(outputString)
                f.Sync()
            }
        }
    }
}

func main() {
    if len(os.Args) < 4 {
        fmt.Printf("Usage: localize input_folder localization_folder output_folder\n")
        os.Exit(1)
    }
    inputFolder := os.Args[1]
    localizationFolder := os.Args[2]
    outputFolder := os.Args[3]

    files, _ := ioutil.ReadDir(localizationFolder)
    for _, f := range files {
        if !f.IsDir() {
            //if filepath.Ext(f.Name()) == ".lcz" {
                fmt.Printf("Localizing %s", f.Name()[len(f.Name()) - len(filepath.Ext(f.Name())) + 1:len(f.Name())])
                data, error := ioutil.ReadFile(fmt.Sprint(localizationFolder + "/" + f.Name()))
                check(error)
                localize := string(data)
                var localization map[string]string
                localization = make(map[string]string)

                var start, end int
                var key, value string
                for {
                    start = strings.Index(localize, "{{")
                    if start == -1 { break; }
                    end = strings.Index(localize, "}}") + len("}}")
                    key = localize[start:end]
                    localize = localize[end:len(localize)]
                    start = strings.Index(localize, "= ") + len("= ")
                    end = strings.Index(localize, "{{")
                    if (end > -1) {
                        value = strings.TrimSpace(localize[start:end])
                        localize = localize[end:len(localize)]
                    } else {
                        value = strings.TrimSpace(localize[start:len(localize)])
                    }
                    localization[key] = value
                }

                Localize(inputFolder, inputFolder, strings.Replace(fmt.Sprint(outputFolder + "/" + f.Name()[len(f.Name()) - len(filepath.Ext(f.Name())) + 1:len(f.Name())]), ".", "-", -1), localization)
                fmt.Printf(" [OK]\n")
            //}
        }
    }

    Localize(inputFolder, inputFolder, outputFolder, nil)
}
