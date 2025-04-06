package main

import (
    "fmt"
    "math/rand"
    "sort"
    "strings"
    "time"
)

func generateWordlist(req Request) []string {
    rand.Seed(time.Now().UnixNano())

    var seeds []string
    switch req.Scope {
    case "institution":
        seeds = append(seeds, splitInput(req.InstitutionInfo.Name)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.Acronyms)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.Departments)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.Locations)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.Years)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.StaffRoles)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.Industry)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.CommonTerms)...)
        seeds = append(seeds, splitInput(req.InstitutionInfo.School)...)
    case "homeowner":
        seeds = append(seeds, splitInput(req.HomeownerInfo.FirstName)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.LastName)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.MiddleName)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.Aliases)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.Usernames)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.DOB)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.Industry)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.JobTitle)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.Anniversaries)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.Keywords)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.Email)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.IDNumber)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.PhoneNumber)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.SSN)...)
        seeds = append(seeds, splitInput(req.HomeownerInfo.PetName)...)
    case "public-wifi":
        seeds = append(seeds, splitInput(req.WifiInfo.LocationName)...)
        seeds = append(seeds, splitInput(req.WifiInfo.OwnerName)...)
        seeds = append(seeds, splitInput(req.WifiInfo.Years)...)
        seeds = append(seeds, splitInput(req.WifiInfo.Keywords)...)
        seeds = append(seeds, splitInput(req.WifiInfo.PhoneNumber)...)
    case "unknown":
        seeds = append(seeds, "guest", "user", "admin", "wifi")
    }

    // Social accounts
    seeds = append(seeds, splitInput(req.SocialAccounts.XHandle)...)
    seeds = append(seeds, splitInput(req.SocialAccounts.Instagram)...)
    seeds = append(seeds, splitInput(req.SocialAccounts.Facebook)...)
    seeds = append(seeds, splitInput(req.SocialAccounts.LinkedIn)...)
    seeds = append(seeds, splitInput(req.SocialAccounts.OtherPlatforms)...)
    seeds = append(seeds, splitInput(req.SocialAccounts.ProfileKeywords)...)
    seeds = append(seeds, splitInput(req.SocialAccounts.Email)...)

    // Extra seeds
    if req.ExtraSeeds != "" {
        seeds = append(seeds, strings.Split(req.ExtraSeeds, "\n")...)
    }

    // Boosters
    if req.AddTopPasswords {
        seeds = append(seeds, "Password123!", "Qwerty2025", "Admin@123", "LetMeIn99", "12345!abc")
    }
    if req.AddKeyboardPatterns {
        seeds = append(seeds, "Qwe123!", "Asd#456", "Zxc$789", "1qaz@WSX", "!QAZxsw2")
    }

    // Clean and dedupe seeds
    var cleanSeeds []string
    seen := make(map[string]bool)
    for _, seed := range seeds {
        seed = strings.TrimSpace(seed)
        if seed != "" && !seen[seed] {
            cleanSeeds = append(cleanSeeds, seed)
            seen[seed] = true
            // Add email fragments
            if strings.Contains(seed, "@") {
                parts := strings.Split(seed, "@")
                if len(parts) > 0 {
                    cleanSeeds = append(cleanSeeds, parts[0]) // e.g., "john.doe"
                }
            }
            // Add phone/ID fragments
            if len(seed) >= 4 {
                cleanSeeds = append(cleanSeeds, seed[len(seed)-4:]) // Last 4 digits
            }
        }
    }

    // Generate variations
    var wordlist []string
    years := []string{"2023", "2024", "2025"}
    if req.PasswordRules.LegacyYears {
        years = append(years, "2020", "2021", "2022")
    }
    specials := "!@#$%^&*+-_"

    for _, seed := range cleanSeeds {
        switch req.PasswordRules.Pattern {
        case "wordnum":
            for _, year := range years {
                wordlist = append(wordlist, fmt.Sprintf("%s%s", seed, year))
                if req.PasswordRules.Uppercase {
                    wordlist = append(wordlist, fmt.Sprintf("%s%s", strings.Title(seed), year))
                }
                if req.PasswordRules.Symbols {
                    special := string(specials[rand.Intn(len(specials))])
                    wordlist = append(wordlist, fmt.Sprintf("%s%s%s", seed, special, year))
                    wordlist = append(wordlist, fmt.Sprintf("%s%s%s", strings.ToUpper(seed), special, year))
                }
            }
        case "wordwordnum":
            for _, s2 := range cleanSeeds {
                if s2 != seed {
                    for _, year := range years {
                        wordlist = append(wordlist, fmt.Sprintf("%s%s%s", seed, s2, year))
                        if req.PasswordRules.Uppercase {
                            wordlist = append(wordlist, fmt.Sprintf("%s%s%s", strings.Title(seed), strings.Title(s2), year))
                        }
                        if req.PasswordRules.Symbols {
                            special := string(specials[rand.Intn(len(specials))])
                            wordlist = append(wordlist, fmt.Sprintf("%s%s%s%s", seed, special, s2, year))
                            wordlist = append(wordlist, fmt.Sprintf("%s%s%s%s", strings.ToUpper(seed), special, strings.ToLower(s2), year))
                        }
                    }
                }
            }
        case "num":
            wordlist = append(wordlist, "12345678!", "87654321#", "11!22@33", "1234$5678", "0000!0000")
            if req.PasswordRules.Uppercase {
                wordlist = append(wordlist, "123ABC!789")
            }
        case "keyboard":
            keyboards := []string{"Qwerty123!", "Asdf#456", "Zxcv$789", "1QAZ!xsw", "Qwe!@#rty", "ZxC!vBn"}
            for _, kb := range keyboards {
                wordlist = append(wordlist, kb)
                if req.PasswordRules.Symbols {
                    wordlist = append(wordlist, kb+string(specials[rand.Intn(len(specials))]))
                }
            }
        default:
            for _, year := range years {
                mixed := mixCase(seed)
                wordlist = append(wordlist, fmt.Sprintf("%s%s", mixed, year))
                if req.PasswordRules.Symbols {
                    special := string(specials[rand.Intn(len(specials))])
                    wordlist = append(wordlist, fmt.Sprintf("%s%s%s", mixed, special, year))
                    wordlist = append(wordlist, fmt.Sprintf("%s%s%s", seed, special, mixed))
                }
                if strings.Contains(seed, "@") {
                    parts := strings.Split(seed, "@")
                    wordlist = append(wordlist, fmt.Sprintf("%s%s", parts[0], year)) // e.g., "john.doe2025"
                }
            }
        }
    }

    // Apply character rules
    allowedChars := buildAllowedChars(req.PasswordRules)
    var filtered []string
    for _, w := range wordlist {
        if len(w) >= req.PasswordRules.MinLength && len(w) <= req.PasswordRules.MaxLength && isValid(w, allowedChars) {
            filtered = append(filtered, w)
        }
        if len(filtered) >= req.OutputSettings.Size {
            break
        }
    }

    if req.OutputSettings.Dedupe {
        filtered = dedupe(filtered)
    }
    if req.OutputSettings.SortByProb {
        sort.Strings(filtered)
    }

    return filtered
}

func mixCase(s string) string {
    chars := []rune(s)
    for i := range chars {
        if rand.Intn(2) == 0 {
            chars[i] = rune(strings.ToUpper(string(chars[i]))[0])
        } else {
            chars[i] = rune(strings.ToLower(string(chars[i]))[0])
        }
    }
    return string(chars)
}

func splitInput(input string) []string {
    if input == "" {
        return []string{}
    }
    return strings.Split(strings.ReplaceAll(input, " ", ""), ",")
}

func buildAllowedChars(rules PasswordRules) string {
    var chars string
    if rules.Lowercase {
        chars += "abcdefghijklmnopqrstuvwxyz"
    }
    if rules.Uppercase {
        chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    }
    if rules.Numbers {
        chars += "0123456789"
    }
    if rules.Symbols {
        chars += "!@#$%^&*+-_"
    }
    if chars == "" {
        chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    }
    return chars
}

func isValid(word, allowed string) bool {
    for _, c := range word {
        if !strings.ContainsRune(allowed, c) {
            return false
        }
    }
    return true
}

func dedupe(list []string) []string {
    seen := make(map[string]bool)
    var result []string
    for _, item := range list {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    return result
}