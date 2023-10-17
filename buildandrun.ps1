$main = "./cmd/main.go"
$file = "binary.exe"

&go build -o $file $main
if ($LASTEXITCODE -eq 0) {
    &./$file
}