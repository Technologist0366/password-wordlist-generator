package main

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/rs/cors"
)

type Request struct {
    Scope              string          `json:"scope"`
    TargetHint         string          `json:"targetHint"`
    InstitutionInfo    InstitutionInfo `json:"institutionInfo,omitempty"`
    HomeownerInfo      HomeownerInfo   `json:"homeownerInfo,omitempty"`
    WifiInfo           WifiInfo        `json:"wifiInfo,omitempty"`
    SocialAccounts     SocialAccounts  `json:"socialAccounts,omitempty"`
    ExtraSeeds         string          `json:"extraSeeds,omitempty"`
    AddTopPasswords    bool            `json:"addTopPasswords"`
    AddKeyboardPatterns bool           `json:"addKeyboardPatterns"`
    PasswordRules      PasswordRules   `json:"passwordRules"`
    OutputSettings     OutputSettings  `json:"outputSettings"`
}

type InstitutionInfo struct {
    Name        string `json:"name"`
    Acronyms    string `json:"acronyms"`
    Departments string `json:"departments"`
    Locations   string `json:"locations"`
    Years       string `json:"years"`
    StaffRoles  string `json:"staffRoles"`
    Industry    string `json:"industry"`
    CommonTerms string `json:"commonTerms"`
    Website     string `json:"website"`
    School      string `json:"school"` // e.g., "Dekut University"
}

type HomeownerInfo struct {
    FirstName    string `json:"firstName"`
    LastName     string `json:"lastName"`
    MiddleName   string `json:"middleName"`
    Aliases      string `json:"aliases"`
    Usernames    string `json:"usernames"`
    DOB          string `json:"dob"`
    Industry     string `json:"industry"`
    JobTitle     string `json:"jobTitle"`
    Anniversaries string `json:"anniversaries"`
    Keywords     string `json:"keywords"`
    Email        string `json:"email"`    // e.g., "john.doe@gmail.com"
    IDNumber     string `json:"idNumber"` // e.g., "123456789"
    PhoneNumber  string `json:"phoneNumber"` // e.g., "254712345678"
    SSN          string `json:"ssn"`      // e.g., "123-45-6789"
    PetName      string `json:"petName"`  // e.g., "Fluffy"
}

type WifiInfo struct {
    LocationName string `json:"locationName"`
    OwnerName    string `json:"ownerName"`
    Years        string `json:"years"`
    Keywords     string `json:"keywords"`
    PhoneNumber  string `json:"phoneNumber"` // e.g., "254712345678"
}

type SocialAccounts struct {
    XHandle        string `json:"xHandle"`
    Instagram      string `json:"instagram"`
    Facebook       string `json:"facebook"`
    LinkedIn       string `json:"linkedin"`
    OtherPlatforms string `json:"otherPlatforms"`
    ProfileKeywords string `json:"profileKeywords"`
    Email          string `json:"email"` // e.g., "john.doe@x.com"
}

type PasswordRules struct {
    MinLength   int    `json:"minLength"`
    MaxLength   int    `json:"maxLength"`
    Lowercase   bool   `json:"lowercase"`
    Uppercase   bool   `json:"uppercase"`
    Numbers     bool   `json:"numbers"`
    Symbols     bool   `json:"symbols"`
    Pattern     string `json:"pattern"`
    LegacyYears bool   `json:"legacyYears"`
}

type OutputSettings struct {
    Size       int    `json:"size"`
    Dedupe     bool   `json:"dedupe"`
    SortByProb bool   `json:"sortByProb"`
    Format     string `json:"format"`
}

func main() {
    r := gin.Default()

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })
    r.Use(func(ctx *gin.Context) {
        c.HandlerFunc(ctx.Writer, ctx.Request)
        if ctx.Request.Method == "OPTIONS" {
            ctx.AbortWithStatus(204)
            return
        }
        ctx.Next()
    })

    r.POST("/generate", func(c *gin.Context) {
        var req Request
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        wordlist := generateWordlist(req)
        previewSize := 5
        if len(wordlist) < previewSize {
            previewSize = len(wordlist)
        }
        c.JSON(http.StatusOK, gin.H{"wordlist": wordlist[:previewSize]})
    })

    r.POST("/download", func(c *gin.Context) {
        var req Request
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        wordlist := generateWordlist(req)
        c.Header("Content-Disposition", "attachment; filename=wordlist.txt")
        c.Header("Content-Type", "text/plain")
        c.Writer.WriteString(strings.Join(wordlist, "\n"))
    })

    r.Run(":8080")
}