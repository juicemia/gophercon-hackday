echo "Compiling..."
GOARCH=arm GOOS=linux go build -o app
echo "Copying..."
scp app pi@$1:/home/pi/app
echo "Running..."
ssh -t pi@$1 ./app
