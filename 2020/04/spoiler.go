(computePartOne(passports))
	fmt.Println(computePartTwo(passports))
}

func computePartOne(passports []string) int {
	numValidPassports := 0

	for _, passport := range passports {
		fieldToValue := parsePassportFieldsAndValues(passport)

		fields := make([]string, 0, len(fieldToValue))
		for field := range fieldToValue {
			fields = append(fields, field)
		}
		sort.Sort(sort.StringSlice(fields))

		if areFieldsValid(fields) {
			numValidPassports++
		}
	}
	return numValidPassports
}

func computePartTwo(passports []string) int {
	numValidPassports := 0

	for _, passport := range passports {
		fieldToValue := parsePassportFieldsAndValues(passport)

		fields := make([]string, 0, len(fieldToValue))
		for field := range fieldToValue {
			fields = append(fields, field)
		}
		sort.Sort(sort.StringSlice(fields))

		if areFieldsValid(fields) && areValuesValid(fieldToValue) {
			numValidPassports++
		}
	}
	return numValidPassports
}

func areValuesValid(fieldToValue map[string]string) bool {
	for field, value := range fieldToValue {
		if field == "byr" && !validateByrValue(value) {
			return false
		}
		if field == "iyr" && !validateIyrValue(value) {
			return false
		}
		if field == "eyr" && !validateEyrValue(value) {
			return false
		}
		if field == "hgt" && !validateHgtValue(value) {
			return false
		}
		if field == "hcl" && !validateHclValue(value) {
			return false
		}
		if field == "ecl" && !validateEclValue(value) {
			return false
		}
		if field == "pid" && !validatePidValue(value) {
			return false
		}
	}
	return true
}

func validateByrValue(v string) bool {
	i, _ := strconv.Atoi(v)

	return i >= 1920 && i <= 2002
}

func validateIyrValue(v string) bool {
	i, _ := strconv.Atoi(v)

	return i >= 2010 && i <= 2020
}

func validateEyrValue(v string) bool {
	i, _ := strconv.Atoi(v)

	return i >= 2020 && i <= 2030
}

func validateHgtValue(v string) bool {
	if strings.HasSuffix(v, "in") {
		splitOnIn := strings.Split(v, "in")
		i, _ := strconv.Atoi(splitOnIn[0])
		return i >= 59 && i <= 76
	}
	if strings.HasSuffix(v, "cm") {
		splitOnIn := strings.Split(v, "cm")
		i, _ := strconv.Atoi(splitOnIn[0])
		return i >= 150 && i <= 193
	}
	return false
}

func validateHclValue(v string) bool {
	valid, _ := regexp.MatchString("^#[0-9a-f]{6}", v)
	return valid
}

func validateEclValue(v string) bool {
	return v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "grn" || v == "hzl" || v == "oth"
}

func validatePidValue(v string) bool {
	valid, _ := regexp.MatchString("^[0-9]{9}", v)
	return valid && len(v) == 9
}

func areFieldsValid(fields []string) bool {
	allFields := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"}
	allFieldsExceptCid := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}

	sort.Sort(sort.StringSlice(allFields))
	sort.Sort(sort.StringSlice(allFieldsExceptCid))

	return reflect.DeepEqual(fields, allFields) || reflect.DeepEqual(fields, allFieldsExceptCid)
}

func getPassports(inputFileName string) ([]string, error) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, fmt.Errorf("Error while opening file %v\n", err)
	}
	defer func() {
		if err = inputFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Error while reading file %v\n", err)
	}

	return strings.Split(string(content), "\n\n"), nil
}

func parsePassportFieldsAndValues(passport string) map[string]string {
	passportFieldsToValues := make(map[string]string)

	splitByNewLine := strings.Split(passport, "\n")
	for _, s := range splitByNewLine {
		fieldValuePairs := strings.Split(s, " ")
		for _, fieldValuePair := range fieldValuePairs {
			splitOnColon := strings.Split(fieldValuePair, ":")
			field, value := splitOnColon[0], splitOnColon[1]
			passportFieldsToValues[field] = value
		}
	}
	return passportFieldsToValues
}