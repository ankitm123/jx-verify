## jx-verify job

Verifies that the job(s) with the given label succeeds and tails the log as it executes

***Aliases**: logs*

### Usage

```
jx-verify job
```

### Synopsis

Verifies that the job(s) with the given label succeeds and tails the log as it executes

### Examples

  # verify the BDD job succeeds
  jx verify job -l app=jx-bdd

### Options

```
  -b, --batch-mode          Runs in batch mode without prompting for user input
  -c, --container string    the name of the container in the job to log (default "job")
  -d, --duration duration   how long to wait for a Job to be active and a Pod to be ready (default 30m0s)
  -h, --help                help for job
      --log-level string    Sets the logging level. If not specified defaults to $JX_LOG_LEVEL
  -n, --namespace string    the namespace where the jobs run. If not specified it will look in: jx-git-operator and jx
      --poll duration       duration between polls for an active Job or Pod (default 1s)
  -l, --selector string     the selector of the job pods
      --verbose             Enables verbose output. The environment variable JX_LOG_LEVEL has precedence over this flag and allows setting the logging level to any value of: panic, fatal, error, warn, info, debug, trace
```

### SEE ALSO

* [jx-verify](jx-verify.md)	 - commands for verifying Jenkins X environments

###### Auto generated by spf13/cobra on 9-Feb-2021