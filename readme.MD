# Flutter assessment test

## Given the following

### Instructions

Included in the zip is a commits.csv file of commits into the default branches of GitHub  
repositories from teams working in an inner source model at Flutter. This is anonymised real data with the following fields:

| Field        | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| `timestamp`  | The unix timestamp of the commit                             |
| `user`       | The GitHub username of the commit author. This data is unreliable and blank indicates the author's identity is unknown. |
| `repository` | The repository name which the commit was pushed to.          |
| `files`      | The number of files changes by the commit                    |
| `additions`  | The number of line additions in this commit                  |
| `deletions`  | The number of deletions in this commit                       |

The file contains 100 days of commits.

## Your Task

Using Golang:

1\. Design and implement an algorithm that will give each repository an activity score.  
2\. Document this algorithm in a markdown file, and any directions required to run your  
implementation of it.  
3\. Use this algorithm to rank the repositories and include the top 10 "most active" repositories by your definition in the documentation.  
4\. Zip up the documentation and your implementation and return it to complete your test.  
The task should take no longer than 1-2 hours. Note that you are welcome to open and explore the CSV data in Excel if you wish, but your ranking algorithm should be implemented in a programming language for automated execution.

### The solution

The algorithm used was the sum of the commit lines per repository. After this, each activity, such as files, additions, and deletions, was summed per repository.

The challenge was to create a code within 2 hours, keeping the code legible and understandable. A big structure was designed to hold the data read from the CSV and as a materialized view to keep the sums in a mapped/indexed structure.

## To run the solution

You must have Docker installed to run the project and internet access to checkout the Docker image with Golang.

Clone this project on your computer and move it to the directory where the project was cloned.

### Windows

If your operational system is Windows, open a Power Shell prompt and execute this command:  

```sh
docker run \-it \--rm \-v ${pwd}:/app \-w /app golang go run .
```

If this is your first time running this Docker image, Docker will show you the following message:  

```sh
Unable to find image 'golang:latest' locally
latest: Pulling from library/golang
```

Wait until the image is completely pulled, and see the results.

### Linux

The same above is true for Linux, but the command is a little different:

```sh
docker run \-it \--rm \-v $(pwd):/app \-w /app golang go run .
```

PS: I created a shell for Linux and a batch file for Windows. You can run these files, too. In the case of Linux, permission to execute should be necessary.
