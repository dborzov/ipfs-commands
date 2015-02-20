# ipfs-commands

Generate IPFS commands + api reference markdown, so we can add it to the website easily. See [go-ipfs/Issue#785](https://github.com/jbenet/go-ipfs/issues/785)

Usage:
```
    go get -v github.com/dborzov/ipfs-commands
    cd $GOPATH/src/github.com/dborzov/ipfs-commands
    ipfs-commands
```

et voila, that results in the filled markdown template from [`TEMPLATE.md`](TEMPLATE.md) in Stdout, like this:

-------

## IPFS command line tool commands


## `ipfs version``: Shows ipfs version information

Returns the current version of ipfs and exits.



## `ipfs get``: Download IPFS objects


Retrieves the object named by <ipfs-path> and stores the data to disk.

By default, the output will be stored at ./<ipfs-path>, but an alternate path
can be specified with '--output=<path>' or '-o=<path>'.

To output a TAR archive instead of unpacked files, use '--archive' or '-a'.

To compress the output with GZIP compression, use '--compress' or '-C'. You
may also specify the level of compression by specifying '-l=<1-9>'.




## `ipfs name``: IPFS namespace (IPNS) tool


IPNS is a PKI namespace, where names are the hashes of public keys, and
the private key enables publishing new (signed) values. In both publish
and resolve, the default value of <name> is your own identity public key.




## `ipfs repo``: Manipulate the IPFS repo


'ipfs repo' is a plumbing command used to manipulate the repo.




## `ipfs update``: Downloads and installs updates for IPFS

ipfs update is a utility command used to check for updates and apply them.



## `ipfs ping``: send echo request packets to IPFS hosts


ipfs ping is a tool to test sending data to other nodes. It finds nodes
via the routing system, send pings, wait for pongs, and print out round-
trip latency information.




## `ipfs bootstrap``: Show or edit the list of bootstrap peers


Running 'ipfs bootstrap' with no arguments will run 'ipfs bootstrap list'.

SECURITY WARNING:

The bootstrap command manipulates the "bootstrap list", which contains
the addresses of bootstrap nodes. These are the *trusted peers* from
which to learn about other peers in the network. Only edit this list
if you understand the risks of adding or removing nodes from this list.





## `ipfs cat``: Show IPFS object data


Retrieves the object named by <ipfs-path> and outputs the data
it contains.




## `ipfs id``: Show IPFS Node ID info


Prints out information about the specified peer,
if no peer is specified, prints out local peers info.

ipfs id supports the format option for output with the following keys:
<id> : the peers id
<aver>: agent version
<pver>: protocol version
<pubkey>: public key




## `ipfs object``: Interact with ipfs objects


'ipfs object' is a plumbing command used to manipulate DAG objects
directly.



## `ipfs refs``: Lists links (references) from an object


Retrieves the object named by <ipfs-path> and displays the link
hashes it contains, with the following format:

  <link base58 hash>

Note: list all refs recursively with -r.




## `ipfs commands``: List all available commands.

Lists all available commands (and subcommands) and exits.



## `ipfs diag``: Generates diagnostic reports





## `ipfs log``: Interact with the daemon log output


'ipfs log' contains utility commands to affect or read the logging
output of a running daemon.




## `ipfs ls``: List links from an object.


Retrieves the object named by <ipfs-path> and displays the links
it contains, with the following format:

  <link base58 hash> <link size in bytes> <link name>




## `ipfs mount``: Mounts IPFS to the filesystem (read-only)


Mount ipfs at a read-only mountpoint on the OS (default: /ipfs and /ipns).
All ipfs objects will be accessible under that directory. Note that the
root will not be listable, as it is virtual. Access known paths directly.

You may have to create /ipfs and /ipfs before using 'ipfs mount':

> sudo mkdir /ipfs /ipns
> sudo chown `whoami` /ipfs /ipns
> ipfs daemon &
> ipfs mount




## `ipfs pin``: Pin (and unpin) objects to local storage





## `ipfs swarm``: swarm inspection tool


ipfs swarm is a tool to manipulate the network swarm. The swarm is the
component that opens, listens for, and maintains connections to other
ipfs peers in the internet.




## `ipfs add``: Add an object to ipfs.


Adds contents of <path> to ipfs. Use -r to add directories.
Note that directories are added recursively, to form the ipfs
MerkleDAG. A smarter partial add with a staging area (like git)
remains to be implemented.




## `ipfs block``: Manipulate raw IPFS blocks


'ipfs block' is a plumbing command used to manipulate raw ipfs blocks.
Reads from stdin or writes to stdout, and <key> is a base58 encoded
multihash.




## `ipfs config``: get and set IPFS config values


ipfs config controls configuration variables. It works like 'git config'.
The configuration values are stored in a config file inside your IPFS
repository.



## `ipfs dht``: Issue commands directly through the DHT
