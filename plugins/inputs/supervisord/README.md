# supervisord Input Plugin

This plugin gathers supervisord processes metrics.

All measurements have hostname and program name tags.

### Configuration
```toml
[[inputs.supervisord]]
  ## Default supervisor RPC host
  # host = "http://127.0.0.1:9001/RPC2"
```

### Processes info

Processes structure based on [supervisord XML-RPC API Documentation](http://supervisord.org/api.html#process-control).

- processInfo:
    - Name          `string`
    - Group         `string`
    - Description   `string`
    - Start         `int64 (unix timestamp)`
    - Stop          `int64 (unix timestamp)`
    - Now           `int64 (unix timestamp)`
    - State         `int16`
    - Statename     `string`
    - Spawnerr      `string`
    - Exitstatus    `int64 (if running is 0)`
    - StdoutLogfile `string`
    - StderrLogfile `string`
    - Pid           `int64`

### Process states

[Process states documentation](http://supervisord.org/subprocess.html#process-states)

| Name      | Code  | Description   |
| :---      |  ---: |          ---: |
|`STOPPED`  |0      |The process has been stopped due to a stop request or has never been started.|
|`STARTING` |10     |The process is starting due to a start request.|
|`RUNNING`  |20     |The process is running.|
|`BACKOFF`  |30     |The process entered the `STARTING` state but subsequently exited too quickly (before the time defined in `startsecs`) to move to the `RUNNING` state.|
|`STOPPING` |40     |The process is stopping due to a stop request.|
|`EXITED`   |100    |The process exited from the `RUNNING` state (expectedly or unexpectedly).|
|`FATAL`    |200    |The process could not be started successfully.|
|`UNKNOWN`  |1000   |The process is in an unknown state (**supervisord** programming error).|
