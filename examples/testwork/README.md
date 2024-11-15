# 修正第一个记忆节点无法关闭示例
关联问题https://github.com/magicsea/behavior3go/issues/17 。  
追踪原版的老bug，记忆节点链关闭的第一个会执行不了onclose，导致memseq节点会错误的记录runningChild不清理。   
错误表现为，因为runningChild没重置，直接从Wait0-2跳入Wait1-2。   
### 修复方法 
start=i+1改start，增加same判断是否调用链修改了（不判断会导致一直在running节点）。   

### 修复前后日志
重现需开打Wait节点的print日志。
- 修复前日志
```
logtest: <nil> job 0-0
wait: Wait0-1 <milliseconds>ms <nil> => 0
wait: Wait0-1 <milliseconds>ms <nil> => 108
wait: Wait0-1 <milliseconds>ms <nil> => 216
wait: Wait0-1 <milliseconds>ms <nil> => 325
wait: Wait0-1 <milliseconds>ms <nil> => 434
wait: Wait0-1 <milliseconds>ms <nil> => 542
wait: Wait0-1 <milliseconds>ms <nil> => 652
wait: Wait0-1 <milliseconds>ms <nil> => 759
wait: Wait0-1 <milliseconds>ms <nil> => 869
wait: Wait0-1 <milliseconds>ms <nil> => 978
wait: Wait0-1 <milliseconds>ms <nil> => 1087
logtest: <nil> job 0-1
logtest: <nil> job 0-2
wait: Wait0-2 <milliseconds>ms <nil> => 0
wait: Wait0-2 <milliseconds>ms <nil> => 109
logtest: <nil> job 1-0
wait: Wait1-1 <milliseconds>ms <nil> => 0
wait: Wait1-1 <milliseconds>ms <nil> => 111
wait: Wait1-1 <milliseconds>ms <nil> => 221
wait: Wait1-1 <milliseconds>ms <nil> => 331
wait: Wait1-1 <milliseconds>ms <nil> => 442
wait: Wait1-1 <milliseconds>ms <nil> => 551
wait: Wait1-1 <milliseconds>ms <nil> => 662
wait: Wait1-1 <milliseconds>ms <nil> => 771
wait: Wait1-1 <milliseconds>ms <nil> => 880
wait: Wait1-1 <milliseconds>ms <nil> => 991
wait: Wait1-1 <milliseconds>ms <nil> => 1098
logtest: <nil> job 1-1
logtest: <nil> job 1-2
wait: Wait1-2 <milliseconds>ms <nil> => 0
logtest: <nil> job 0-0
wait: Wait0-1 <milliseconds>ms <nil> => 0
wait: Wait0-1 <milliseconds>ms <nil> => 109
wait: Wait0-1 <milliseconds>ms <nil> => 219
wait: Wait0-1 <milliseconds>ms <nil> => 329
wait: Wait0-1 <milliseconds>ms <nil> => 438
wait: Wait0-1 <milliseconds>ms <nil> => 545
wait: Wait0-1 <milliseconds>ms <nil> => 652
wait: Wait0-1 <milliseconds>ms <nil> => 760
wait: Wait0-1 <milliseconds>ms <nil> => 869
wait: Wait0-1 <milliseconds>ms <nil> => 976
wait: Wait0-1 <milliseconds>ms <nil> => 1084
logtest: <nil> job 0-1
logtest: <nil> job 0-2
wait: Wait0-2 <milliseconds>ms <nil> => 0
wait: Wait0-2 <milliseconds>ms <nil> => 108
wait: Wait1-2 <milliseconds>ms <nil> => 0<--------------------错误行,应该进入：job 1-0
wait: Wait1-2 <milliseconds>ms <nil> => 107
logtest: <nil> job 1-0
wait: Wait1-1 <milliseconds>ms <nil> => 0
wait: Wait1-1 <milliseconds>ms <nil> => 108
wait: Wait1-1 <milliseconds>ms <nil> => 216


```

- 修复后日志
```
logtest: <nil> job 0-0
wait: Wait0-1 <milliseconds>ms <nil> => 0
wait: Wait0-1 <milliseconds>ms <nil> => 110
wait: Wait0-1 <milliseconds>ms <nil> => 217
wait: Wait0-1 <milliseconds>ms <nil> => 327
wait: Wait0-1 <milliseconds>ms <nil> => 437
wait: Wait0-1 <milliseconds>ms <nil> => 545
wait: Wait0-1 <milliseconds>ms <nil> => 653
wait: Wait0-1 <milliseconds>ms <nil> => 763
wait: Wait0-1 <milliseconds>ms <nil> => 872
wait: Wait0-1 <milliseconds>ms <nil> => 980
wait: Wait0-1 <milliseconds>ms <nil> => 1088
logtest: <nil> job 0-1
logtest: <nil> job 0-2
wait: Wait0-2 <milliseconds>ms <nil> => 0
wait: Wait0-2 <milliseconds>ms <nil> => 109
logtest: <nil> job 1-0
wait: Wait1-1 <milliseconds>ms <nil> => 0
wait: Wait1-1 <milliseconds>ms <nil> => 109
wait: Wait1-1 <milliseconds>ms <nil> => 219
wait: Wait1-1 <milliseconds>ms <nil> => 328
wait: Wait1-1 <milliseconds>ms <nil> => 436
wait: Wait1-1 <milliseconds>ms <nil> => 545
wait: Wait1-1 <milliseconds>ms <nil> => 652
wait: Wait1-1 <milliseconds>ms <nil> => 759
wait: Wait1-1 <milliseconds>ms <nil> => 868
wait: Wait1-1 <milliseconds>ms <nil> => 976
wait: Wait1-1 <milliseconds>ms <nil> => 1085
logtest: <nil> job 1-1
logtest: <nil> job 1-2
wait: Wait1-2 <milliseconds>ms <nil> => 0
logtest: <nil> job 0-0
wait: Wait0-1 <milliseconds>ms <nil> => 753
wait: Wait0-1 <milliseconds>ms <nil> => 861
wait: Wait0-1 <milliseconds>ms <nil> => 967
wait: Wait0-1 <milliseconds>ms <nil> => 1075
logtest: <nil> job 0-1
logtest: <nil> job 0-2
wait: Wait0-2 <milliseconds>ms <nil> => 0
wait: Wait0-2 <milliseconds>ms <nil> => 109
<---------------------------------------------------------------------正确切换到1-0
logtest: <nil> job 1-0
wait: Wait1-1 <milliseconds>ms <nil> => 0
wait: Wait1-1 <milliseconds>ms <nil> => 108
wait: Wait1-1 <milliseconds>ms <nil> => 217
wait: Wait1-1 <milliseconds>ms <nil> => 325
wait: Wait1-1 <milliseconds>ms <nil> => 433

```