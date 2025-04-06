# Password Wordlist Tool (2025 Edition)

A slick, customizable tool to generate realistic password wordlists for penetration testing, security research, or just geekin' out. Built with a Go backend and React frontend, it crafts passwords based on target scopes—think institutions, homeowners, or public Wi-Fi—using personal data like emails, IDs, phone numbers, pets, and more. Outputs a preview of the top 5 and a full downloadable `.txt` file.

## Features
- **Target Scopes**: Institution (e.g., "Dekut@ICT2025"), Homeowner (e.g., "John!Fluffy2023"), Public Wi-Fi (e.g., "Kapengting#24"), Unknown (generic sweep).
- **Real-World Patterns**: Mixes uppercase, lowercase, numbers, special chars (`!@#$%^&*+-_`), keyboard routes (e.g., "Qwerty123!"), and multi-word combos.
- **Personal Data**: Uses emails (e.g., "john.doe2025"), IDs, phone numbers (last 4 digits), SSNs, pet names, schools, etc.
- **Customizable**: Set min/max length, character types, patterns ("wordnum", "wordwordnum", "keyboard"), and output size.
- **Output**: Preview top 5 in UI, download full list as `wordlist.txt`.

## Prerequisites
- **Go**: 1.16+ (for backend)
- **Node.js**: 16+ (for frontend, with npm)
- **Git**: Optional, for cloning

## Setup

### Clone the Repo (if applicable)
```bash
git clone <your-repo-url>
cd password-wordlist-tool
