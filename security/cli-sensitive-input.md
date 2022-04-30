## Pass sensitive data to your CLI

### Invalid approach

You shouldn't use, `echo` command , environment variables, or CLI flags to pass a sensitive data to any CLI.
The problem here is that it may be captured by shell history or process listings. So some 3rd party program can list running processes and steal sensitive data.
<!--more-->

The `curl` accepts the sensitive data `username:password` in CLI flags, but they try to minimize the risk by removing input arguments in the runtime, see:

```cpp
void cleanarg(char *str)
{
#ifdef HAVE_WRITABLE_ARGV
  /* now that GetStr has copied the contents of nextarg, wipe the next
   * argument out so that the username:password isn't displayed in the
   * system process list */
  if(str) {
    size_t len = strlen(str);
    memset(str, ' ', len);
  }
#else
  (void)str;
#endif
}
```

_source: https://github.com/curl/curl/blob/3085ccfae996bb0fa606d2c7bc6783dc15d76a30/src/tool_paramhlp.c#L112-L125_

#### Try to modify args from Go runtime

In general, bad idea.

_source: https://stackoverflow.com/a/14943149_

### Valid approach

You should read the sensitive data from file, or from `stdin`. Examples:

```bash
# pipe file content
cat ~/my_password.txt | docker login --username foo --password-stdin

# start interactive setup
$ vault login
Token (will be hidden):

# start interactive setup
$ gh auth login

# authenticate against github.com by reading the token from a file
$ gh auth login --with-token < mytoken.txt
```
