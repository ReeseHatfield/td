if ! go build -o td src/*.go; then
    echo "Could not build td, is go installed properly?"
else
    sudo mv td /bin
fi