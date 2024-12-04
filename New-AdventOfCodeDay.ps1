# Yes this script is made by chatgpt
param(
    [string]$TemplateName,
    [int]$Year,
    [int]$Day
)

# Validate Parameters
if (-not $TemplateName) {
    Write-Error "You must specify a template name."
    exit 1
}
if (-not $Year) {
    Write-Error "You must specify the year."
    exit 1
}
if (-not $Day) {
    Write-Error "You must specify the day."
    exit 1
}

# Define paths
$repoRoot = (Get-Location).Path
$templateFolder = Join-Path -Path $repoRoot -ChildPath ".templates\$TemplateName"
$newDayFolder = Join-Path -Path $repoRoot -ChildPath "$Year\Day-$Day"

# Check template existence
if (-not (Test-Path -Path $templateFolder)) {
    Write-Error "Template '$TemplateName' does not exist in the '.templates' folder."
    exit 1
}

# Create day folder
if (-not (Test-Path -Path $newDayFolder)) {
    New-Item -ItemType Directory -Path $newDayFolder | Out-Null
} else {
    Write-Error "The folder for '$Year/Day-$Day' already exists."
    exit 1
}

# Copy template contents
Copy-Item -Path "$templateFolder\*" -Destination $newDayFolder -Recurse -Force

# Change to repository root and run git commands
Set-Location -Path $repoRoot

# Stage changes
git add "$Year/Day-$Day"

# Commit changes
$commitMessage = "init day $Day"
git commit -m $commitMessage

# Find all input.txt files and assume unchanged
$inputFiles = Get-ChildItem -Path $newDayFolder -Recurse -Filter "input.txt"
if ($inputFiles.Count -gt 0) {
    foreach ($inputFile in $inputFiles) {
        git update-index --assume-unchanged $inputFile.FullName
        Write-Host "Marked '$($inputFile.FullName)' as assume-unchanged."
    }
} else {
    Write-Host "No 'input.txt' files found in '$newDayFolder'."
}

Write-Host "Successfully initialized day $Day under $Year with template '$TemplateName'."
