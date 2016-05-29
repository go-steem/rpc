# Voting Monitor

The executable built from this package will connect to `steemd`
and it will start printing information on vote operations as they
are happening and it will keep doing so forever. Well, until interrupted.

```
$ ./monitor_voting -rpc_endpoint="ws://$(docker-machine ip default):8090"
2016/05/29 10:42:56 ---> Dial("ws://192.168.99.100:8090")
2016/05/29 10:42:56 ---> GetConfig()
2016/05/29 10:42:56 ---> Entering the block processing loop (last block = 1866869)
@easteagle13 voted for @easteagle13/another-article-discussing-some-inherent-flaws-of-the-dao
@easteagle13 voted for @easteagle13/to-your-loss-of-a-friend-my-condolences-and-other-thoughts
@yefet voted for @alexgr/planning-for-long-term-success-of-steemit-identifying-areas-of-improvement
@dke voted for @steemrollin/steem-meme
```
