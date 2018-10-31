# Copy
Copy file src to dest, It doesn't matter if src is a directory or a file, exclude is fuzzy string match
# Example Usage

```go
err := Copy("src","dest",[]string{"excludeFile","excludeDir","excludeOther"})
if err != nil{
	fmt.Errorf("error copying file: %v\n", err)
}
```

# License

[MIT License](http://zealzhangz.mit-license.org/)