# Minio Admin Complete Guide [![Slack](https://slack.minio.io/slack?type=svg)](https://slack.minio.io)

Minio Client (mc) provides `admin` sub-command to perform administrative tasks on your Minio deployments.

```sh
service      Stop, restart or get status of Minio server
info         Display Minio server information
users        Manage regular users
policies     Manage canned policies
credentials  Change Admin server access and secret keys
config       Manage configuration file
heal         Heal disks, buckets and objects on Minio server
```

## 1.  Download Minio Client
### Docker Stable
```
docker pull minio/mc
docker run minio/mc admin info play
```

### Docker Edge
```
docker pull minio/mc:edge
docker run minio/mc:edge admin info play
```

### Homebrew (macOS)
Install mc packages using [Homebrew](http://brew.sh/)

```sh
brew install minio/stable/mc
mc --help
```

### Binary Download (GNU/Linux)
| Platform | Architecture | URL |
| ---------- | -------- |------|
|GNU/Linux|64-bit Intel|https://dl.minio.io/client/mc/release/linux-amd64/mc |

```sh
chmod +x mc
./mc --help
```

### Binary Download (Microsoft Windows)
| Platform | Architecture | URL |
| ---------- | -------- |------|
|Microsoft Windows|64-bit Intel|https://dl.minio.io/client/mc/release/windows-amd64/mc.exe |

```sh
mc.exe --help
```

### Install from Source
Source installation is intended only for developers and advanced users. `mc update` command does not support update notifications for source based installations. Please download official releases from https://minio.io/downloads/#minio-client.

If you do not have a working Golang environment, please follow [How to install Golang](https://docs.minio.io/docs/how-to-install-golang).

```sh
go get -d github.com/minio/mc
cd ${GOPATH}/src/github.com/minio/mc
make
```

## 2. Run Minio Client

### GNU/Linux

```sh
chmod +x mc
./mc --help
```

### macOS

```sh
chmod 755 mc
./mc --help
```

### Microsoft Windows

```sh
mc.exe --help
```

## 3. Add a Minio Storage Service
Minio server displays URL, access and secret keys.

#### Usage

```sh
mc config host add <ALIAS> <YOUR-MINIO-ENDPOINT> <YOUR-ACCESS-KEY> <YOUR-SECRET-KEY>
```

Alias is simply a short name to your Minio service. Minio end-point, access and secret keys are supplied by your Minio service. Admin API uses "S3v4" signature and cannot be changed.

```sh
mc config host add minio http://192.168.1.51:9000 BKIKJAA5BMMU2RHO6IBB V7f1CwQqAcwo80UEIJEjc5gVQUSSx5ohQ9GSrr12
```

## 4. Test Your Setup

*Example:*

Get Minio server information for the configured alias `minio`

```sh
mc admin info minio

●  192.168.1.51:9000
   Uptime : online since 1 day ago
  Version : 2018-05-28T04:31:38Z
   Region :
 SQS ARNs : <none>
    Stats : Incoming 82GiB, Outgoing 28GiB
  Storage : Used 7.4GiB
```

## 5. Everyday Use
You may add shell aliases for info, healing.

```sh
alias minfo='mc admin info'
alias mheal='mc admin heal'
```

## 6. Global Options

### Option [--debug]
Debug option enables debug output to console.

*Example: Display verbose debug output for `info` command.*

```sh
mc admin --debug info minio
mc: <DEBUG> GET /minio/admin/v1/info HTTP/1.1
Host: 192.168.1.51:9000
User-Agent: Minio (linux; amd64) madmin-go/0.0.1 mc/2018-05-23T23:43:34Z
Authorization: AWS4-HMAC-SHA256 Credential=**REDACTED**/20180530/us-east-1/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=**REDACTED**
X-Amz-Content-Sha256: UNSIGNED-PAYLOAD
X-Amz-Date: 20180530T001808Z
Accept-Encoding: gzip

mc: <DEBUG> HTTP/1.1 200 OK
Transfer-Encoding: chunked
Accept-Ranges: bytes
Content-Security-Policy: block-all-mixed-content
Content-Type: application/json
Date: Wed, 30 May 2018 00:18:08 GMT
Server: Minio/DEVELOPMENT.2018-05-28T04-31-38Z (linux; amd64)
Vary: Origin
X-Amz-Request-Id: 1533440573A63034
X-Xss-Protection: "1; mode=block"

mc: <DEBUG> Response Time:  140.70112ms

●  192.168.1.51:9000
   Uptime : online since 1 day ago
  Version : 2018-05-28T04:31:38Z
   Region :
 SQS ARNs : <none>
    Stats : Incoming 82GiB, Outgoing 28GiB
  Storage : Used 7.4GiB
```

### Option [--json]
JSON option enables parseable output in JSON format.

*Example: Minio server information.*

```sh
mc admin --json info minio
{
  "status": "success",
  "service": "on",
  "address": "192.168.1.51:9000",
  "error": "",
  "storage": {
    "used": 7979370172,
    "backend": {
      "backendType": "FS"
    }
  },
  "network": {
    "transferred": 90473434722,
    "received": 30662519192
  },
  "server": {
    "uptime": 157467244813288,
    "version": "2018-05-28T04:31:38Z",
    "commitID": "7d8c5ffb13334f4aec20a35bd2575bd7c740fb7a",
    "region": "",
    "sqsARN": []
  }
}
```

### Option [--no-color]
This option disables the color theme. It is useful for dumb terminals.

### Option [--quiet]
Quiet option suppress chatty console output.

### Option [--config-folder]
Use this option to set a custom config path.

### Option [ --insecure]
Skip SSL certificate verification.

## 7. Commands

|   |
|:---|
|[**service** - Start, stop or get the status of Minio server](#service) |
|[**info** - Display Minio server information](#info) |
|[**users** - Manage users](#users) |
|[**policies** - Manage policies](#policies) |
|[**credentials** - Change **admin** server access and secret keys](#credentials) |
|[**config** - Manage server configuration file](#config)|
|[**heal** - Heal disks, buckets and objects on Minio server](#heal) |

<a name="service"></a>
### Command `service` - Stop, restart or get status of Minio server
`service` command provides a way to restart, stop one or get the status of Minio servers (distributed cluster)

```sh
NAME:
  mc admin service - Stop, restart or get status of Minio server

FLAGS:
  --help, -h                       Show help.

COMMANDS:
  status   Get the status of Minio server
  restart  Restart Minio server
  stop     Stop Minio server
```

*Example: Display service uptime for Minio server.*

```sh
mc admin service status play
Uptime: 1 days 19 hours 57 minutes 39 seconds.
```

*Example: Restart remote minio service.*

NOTE: `restart` and `stop` sub-commands are disruptive operations for your Minio service, any on-going API operations will be forcibly canceled. So, it should be used only under certain circumstances. Please use it with caution.

```sh
mc admin service restart play
Restarted `play` successfully.
```

<a name="info"></a>
### Command `info` - Display Minio server information
`info` command displays server information of one or many Minio servers (under distributed cluster)

```sh
NAME:
  mc admin info - Get Minio server information

FLAGS:
  --help, -h                       Show help.
```

*Example: Display Minio server information.*

```sh
mc admin info play
●  play.minio.io:9000
   Uptime : online since 1 day ago
  Version : 2018-05-28T04:31:38Z
   Region :
 SQS ARNs : <none>
    Stats : Incoming 82GiB, Outgoing 28GiB
  Storage : Used 8.2GiB
```

<a name="policies"></a>
### Command `policies` - Manage user policies
`policies` command to add, remove, list policies on Minio server.

```sh
NAME:
  mc admin policies - Manage policies

FLAGS:
  --help, -h                       Show help.

COMMANDS:
  add      Add new policies
  remove   Remove policies
  list     List all policies
```

*Example: Add a new policy 'newpolicy' on Minio, with policy from /tmp/newpolicy.json.*

```sh
mc admin policies add myminio/ newpolicy /tmp/newpolicy.json
```

*Example: Remove policy 'newpolicy' on Minio.*

```sh
mc admin policies remove myminio/ newpolicy
```

*Example: List all policies on Minio.*

```sh
mc admin policies list --json myminio/
{"status":"success","policy":"newpolicy"}
```

<a name="users"></a>
### Command `users` - Manage regular users
`users` command to add, remove, enable, disable, list users on Minio server.

```sh
NAME:
  mc admin users - Manage regular users

FLAGS:
  --help, -h                       Show help.

COMMANDS:
  add      Add new users
  policy   Set policy for user
  disable  Disable users
  enable   Enable users
  remove   Remove users
  list     List all users
```

*Example: Add a new user 'newuser' on Minio, with 'newpolicy' policy.*

```sh
mc admin users add myminio/ newuser newuser123 newpolicy
```

*Example: Change policy for a user 'newuser' on Minio to 'writeonly' policy.*

```sh
mc admin users policy myminio/ newuser writeonly
```

*Example: Disable a user 'newuser' on Minio.*

```sh
mc admin users disable myminio/ newuser
```

*Example: Enable a user 'newuser' on Minio.*

```sh
mc admin users enable myminio/ newuser
```

*Example: Remove user 'newuser' on Minio.*

```sh
mc admin users remove myminio/ newuser
```

*Example: List all users on Minio.*

```sh
mc admin users list --json myminio/
{"status":"success","accessKey":"newuser","userStatus":"enabled"}
```

<a name="credentials"></a>
### Command `credentials` - Change server **admin** access and secret keys
`credentials` command to set new **admin** credentials of a Minio server.

```sh
NAME:
  mc admin credentials - Change server **admin** access and secret keys

FLAGS:
  --help, -h                       Show help.
```

*Example: Set new admin credentials of a Minio server represented by its alias 'myminio'.*

```sh
mc admin credentials myminio/ minio minio123
```

<a name="config"></a>
### Command `config` - Manage server configuration
`config` command to manage Minio server configuration.

```sh
NAME:
  mc admin config - Manage configuration file

USAGE:
  mc admin config COMMAND [COMMAND FLAGS | -h] [ARGUMENTS...]

COMMANDS:
  get  Get config of a Minio server/cluster.
  set  Set new config file to a Minio server/cluster.

FLAGS:
  --help, -h                       Show help.
```

*Example: Get server configuration of a Minio server/cluster.*

```sh
mc admin config get myminio > /tmp/my-serverconfig
```

*Example: Set server configuration of a Minio server/cluster.*

```sh
mc admin config set myminio < /tmp/my-serverconfig
```

<a name="heal"></a>
### Command `heal` - Heal disks, buckets and objects on Minio server
`heal` command heals disks, missing buckets, objects on Minio server. NOTE: This command is only applicable for Minio erasure coded setup (standalone and distributed).

```sh
NAME:
  mc admin heal - Heal disks, buckets and objects on Minio server

FLAGS:
  --recursive, -r                  Heal recursively
  --dry-run, -n                    Only inspect data, but do not mutate
  --force-start, -f                Force start a new heal sequence
  --help, -h                       Show help.
```

*Example: Heal Minio cluster after replacing a fresh disk, recursively heal all buckets and objects, where 'myminio' is the Minio server alias.*

```sh
mc admin heal -r myminio
```

*Example: Heal Minio cluster on a specific bucket recursively, where 'myminio' is the Minio server alias.*

```sh
mc admin heal -r myminio/mybucket
```

*Example: Heal Minio cluster on a specific object prefix recursively, where 'myminio' is the Minio server alias.*

```sh
mc admin heal -r myminio/mybucket/myobjectprefix
```
