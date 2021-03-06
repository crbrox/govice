/**
 * @license
 * Copyright 2017 Telefónica Investigación y Desarrollo, S.A.U
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/Telefonica/govice"
)

type config struct {
	Address  string `json:"address" env:"ADDRESS"`
	BasePath string `json:"basePath" env:"BASE_PATH"`
	LogLevel string `json:"logLevel" env:"LOG_LEVEL"`
}

func main() {
	var cfg config
	if err := govice.GetConfig("config.json", &cfg); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cfg)

	validator := govice.NewValidator()
	if err := validator.LoadSchemas("schemas"); err != nil {
		panic(err)
	}
	if err := validator.ValidateConfig("config", &cfg); err != nil {
		panic(err)
	}
	fmt.Println("Configuration validated successfully")
}
