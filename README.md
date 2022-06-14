# Backup Utility

**Disclamer:** This is still a work in progress under MIT Licence.

This project aims to create a Linux compatible wrapper to run rsync/tar commands 
to create, access, test and destroy filesystem backups.

If you are looking for entrepise/production ready solution this may not be
the best option. There are some filesystem (ex: zfs) that work better than this.

## Why use linux commands?

The are several reasons:

1. The most important reason, my only use case is to backup multiple servers of 
mine more easily. Later, if there is interest, I may be available to explore 
other possibilities.
2. `rsync` is used for every system administrator and it's very well tested. 
3. `scp` it's not an option because it's very limited.
4. `tar` offer compression in the destination and incremental snapshots.

## Instalation

TBD

## Usage

This steps will be automated in the future. For now, before you create a
backup flow you must:
1. Create an ssh key for this program `ssh-keygen`.
   + Never share you public or private key for any reason. 
2. Copy this now key to your destination server
    + Option 1 (Manually):
      + Manually copy you new generated id_rsa.pub to this file 
      `~/.ssh/authorized_keys` on the destination server
    + Option 2 (ssh-copy-id):
      + TBD

After verify that the source  

TBD

## Further work

(The order will change)
+ [ ] Push flow
+ [ ] Pull flow
+ [ ] Text file based configuration
+ [ ] Internal DB configuration
+ [ ] Encrypt destination file

## Security advices
1. Never use the `root` account.
2. If the destination server is yours, you should not share the account to 
administer the server with the one used by this program. You can create a 
new user `useradd <newuser>` and add your user to the same user group as that 
new user `sudo usermod -a -G <newuser> $USER`.

## Licence

MIT License

Copyright (c) 2022 MauroMaia

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
