# go_crawl
A simple web site crawler written in Go. It takes a starting URL, and will crawl all links it finds in the same domain as the starting URL. It will output a list of all web pages it finds as well as what links to the same domain are on each page. It will also list all images, stylesheets, and javascript files referenced by the page, even if they are from a different domain.

I'm still learning Go, so some things are obviously not going to be the proper way of doing things.

## Building

This code requires the ```net/html``` package that is not included in most Linux distribution's golang package. You can install it by running:

    go get golang.org/x/net/html

Once that package is installed you can build the executable by running:

    go build

## Running

Once it's built, run:

    ./go_crawl -u http://example.com/

## Testing

The ```test_data``` directory contains HTML pages needed for the test cases. The test cases assume that these pages will be accessible at ```http://localhost:8000```. The simplest way to accomplish this is to run the following in a separate terminal:

    cd test_data
    python -m SimpleHTTPServer

Once the test data is accessible, run the tests:

    go test
