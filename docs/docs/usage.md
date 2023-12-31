---
sidebar_position: 2
---

# Usage

## List

Lists all buckets avaliable in your bucket directory.

```
$ autoexec list
  your
  buckets
  here
```

## Create

To create a bucket, use the create command. Use any name for the bucket and a txt file will be created with the name. 

```
$ autoexec create <bucket name>
```

## Add

To add a program to your bucket, use the add command. 

```
$ autoexec add <bucket>
```

Another prompt will appear, asking for the path to the program. Type in the path. The program will automatically verify if the path exists, meaning there is a file of some sort at the location. 

```
$ Enter file path: C:/Your/Path goes/Here.exe
```

If you want to add to a bucket that does not exist, it will create a new bucket with the name you give the first command.

## Read

This command just reads the paths that are currently in the bucket.

```
$ autoexec read <bucket>
  C:/Your/Paths
  C:/Printed/Here
```

## Run

This command runs all the files in your bucket. All it does is launch what ever is at the path so be careful that they are safe to launch as such. 

```
$ autoexec run <bucket>
  Running C:/Your/Paths
  Running C:/Printed/Here
```

## Help

As you would expect, the help command prints out a description of all the avaliable commands.

```
$ autoexec help
  USAGE:
      list            - List avaliable buckets
      create <bucket> - Create a new bucket
      add    <bucket> - Add a path to a bucket
      read   <bucket> - Read the contents of a bucket
      run    <bucket> - Run the contents of a bucket
      help            - Display this message
```