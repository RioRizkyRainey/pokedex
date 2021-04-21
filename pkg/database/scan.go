package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go/ast"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//Field struct Field
type Field struct {
	Name       string
	ColumnName string
	Type       reflect.Type
}

//Scan func to scan
func Scan(rows *sql.Rows, destination interface{}) error {
	defer func() {
		spew.Dump("closed")
		rows.Close()
	}()
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))

	switch destination.(type) {
	case *int, *int8, *int16, *int32, *int64,
		*uint, *uint8, *uint16, *uint32, *uint64, *uintptr,
		*float32, *float64,
		*bool, *string, *time.Time,
		*sql.NullInt32, *sql.NullInt64, *sql.NullFloat64,
		*sql.NullBool, *sql.NullString, *sql.NullTime:
		for rows.Next() {
			rows.Scan(destination)
		}
	default:
		spew.Dump("--> default")
		rv := reflect.Indirect(reflect.ValueOf(destination))

		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		spew.Dump("--> " + rv.Kind().String())
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			modelType := rv.Type().Elem()
			if modelType.Kind() == reflect.Ptr {
				modelType = modelType.Elem()
			}

			fields := make(map[string]*Field, 0)

			for i := 0; i < modelType.NumField(); i++ {
				if fieldStruct := modelType.Field(i); ast.IsExported(fieldStruct.Name) {
					fieldName := fieldStruct.Name
					columnName := ToSnakeCase(fieldName)
					fields[columnName] = &Field{
						Name:       fieldName,
						ColumnName: columnName,
						Type:       fieldStruct.Type,
					}
				}
			}

			for rows.Next() {
				for idx, column := range columns {
					if field := fields[column]; field != nil {
						values[idx] = reflect.New(field.Type).Interface()
					}
				}

				err := rows.Scan(values...)

				if err != nil {
					spew.Dump("--> #1")
					fmt.Println(err.Error())
				}

				modelValue := reflect.New(modelType)

				for idx, column := range columns {
					if field := fields[column]; field != nil {
						modelValue.Elem().FieldByName(field.Name).Set(reflect.ValueOf(values[idx]).Elem())
					}
				}

				rv.Set(reflect.Append(rv, modelValue))
			}
		case reflect.Struct:
			modelType := rv.Type()

			fields := make(map[string]*Field, 0)

			for i := 0; i < modelType.NumField(); i++ {
				if fieldStruct := modelType.Field(i); ast.IsExported(fieldStruct.Name) {
					fieldName := fieldStruct.Name
					columnName := ToSnakeCase(fieldName)
					fields[columnName] = &Field{
						Name:       fieldName,
						ColumnName: columnName,
						Type:       fieldStruct.Type,
					}
				}
			}

			if rows.Next() {
				for idx, column := range columns {
					if field := fields[column]; field != nil {
						values[idx] = reflect.New(field.Type).Interface()
					}
				}

				err := rows.Scan(values...)

				if err != nil {
					spew.Dump("--> #2")
					fmt.Println(err.Error())
					return err
				}

				for idx, column := range columns {
					if field := fields[column]; field != nil {
						rv.FieldByName(field.Name).Set(reflect.ValueOf(values[idx]).Elem())
					}
				}
				return nil
			}
			return errors.New("Error")
		}
	}
	return nil
}

//ToSnakeCase make field name to snake case
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
