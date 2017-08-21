# short_pwd - Go program to get a shortened cwd for shell prompts

## Example

Call the binary in your shell config, eg. in bash by including `$(short_pwd)` in PS1 declaration,
and then ensuring it's called with each prompt render by adding a PROMPT_COMMAND env var:

~~~bash
PS1="<prompt definition>$(short_pwd)$ "
PROMPT_COMMAND='PS1="<PS1 from above>"'
~~~
