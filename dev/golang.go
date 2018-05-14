package dev


import (
  "io"
  "os"
  "log"
  "strings"
)


func Gopack(package_name string) error {
  log.Println("Gopack")

  folderstem := "./" + package_name
  filestem := folderstem + "/" + package_name

  err := os.Mkdir(folderstem, os.ModePerm)

  fbasic, err := os.Create(filestem + ".go")
  if err != nil { return err }
  defer fbasic.Close()

  ftest, err := os.Create(filestem + "_test.go")
  if err != nil { return err }
  defer ftest.Close()

  // write to the files
  basic_string := "package " + package_name + "\n\nimport (\n"
  testing_string := basic_string + "  \"testing\"\n)\n\n"
  testing_string += "func TestBasic(t *testing.T) {\n\n}"
  basic_string += "  \n)\n\n"

  _, err = io.Copy(fbasic, strings.NewReader(basic_string))
  if err != nil { return err }

  _, err = io.Copy(ftest, strings.NewReader(testing_string))
  if err != nil { return err }

  return nil
}
