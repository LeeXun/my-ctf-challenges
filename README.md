# My CTF Challenges

This repository contains my CTF challenges. Write up is in the README.md and source code is seperated by different CTF.
Enjoy! :)


## Balsn CTF 2019

### Gopher party (O口o)!!!(O口o)!!!(O口o)!!!

- Solved Ratio: 5/720

#### Solution

0. Well, I spent lots of time on building the beautiful frontend page. Hope you like it :)
1. In [go.mod](balsn-ctf-2019/gopher-party/go.mod), which can discover it is go1.13. But go scheduler was [released in go1.1](http://morsmachine.dk/go-scheduler) so this won't be a problem.
3. Looking into [main.go](balsn-ctf-2019/gopher-party/main.go), which can find my hint about setting `runtime.GOMAXPROCS(1)` and `t2.nano`.
4. Global search "flag" and ignore the vendors, you may find there's only one target in [register.go](balsn-ctf-2019/gopher-party/controller/register.go).
5. Simple as it goes: you should be the chosen to get the flag.
6. However, after looking for where the sp is defined in [store/store.go](balsn-ctf-2019/gopher-party/store/store.go). You can discover that there's a goroutine keeps changing `the chosen` in every 2ms. 
7. Now the goal is simple: <b>try not to be preempted by other goroutines</b>.
7. This is a challenge of how to avoid racing by the goroutine and other participants' requests.
8. Actually, goroutine would not be preemted by others. Instead, <b>it yields itself and switches context to other goroutines in a magic function which is injected by go compiler: morestack().</b>
9. <b>In conclusion:</b>
    1. Don't let unbuffered channels get blocked.
    1. Avoid making any RPC, like redis or google api. Which will cause netpolling.
    1. Avoid accidentally making the stack size of goroutine being too large.
    1. Avoid sleeping.
10. And there's one more thing that most of participants missed, <b>don't make the running time of your goroutine exceed 10ms</b>.
11. There's a sysmon which is running on a different OS thread then user goroutines. Which means, even if you avoid all the conditions above, OS may still swap your OS thread out for sysmon. When the sysmon detects that your goroutine has exceeded 10m, sysmon would set a magic value (0xfffffade) to your goroutine's stack edge address. The next time your goroutine is up, it would be cheated as if it's stack size is not enough. Then it would yield to next the goroutine in the morestack function. As a result, this goroutine won't be the chosen one.
12. That is, you should find a language string in Accept-Language format that won't take too long in sha256 to complete the payload.
13. Even though this is not a RCE exploit, most of the golang writers who don't REALLY understand how goroutine works, will easily encounter some unexpected issues and finally lead to system impairment. ^(Owo)-o _(XoX)rz

#### Questions

1. Why can I get the flag sometimes when I am running the code at my local machine?
    - The connection between your golang server and redis is too fast. This is a trap.

#### Exploit

```bash
# Google account name is name1
# name1 shouldn't contain any char in the range of https://github.com/LeeXun/my-ctf-challenges/blob/master/balsn-ctf-2019/gopher-party/config/config.go#L8
TOKEN="access_token"
HOST="http://localhost:8000/register"

function exploit()
{
  curl "${HOST}" \
  -H 'accept-language: hu' \
  -H 'content-type: application/x-www-form-urlencoded' \
  --data "name=different_from_name1&access_token=${TOKEN}&interest=AH!&age=-1&praise=&prove="
}

exploit
```

---
Copyright © Lee Xun, 2019
