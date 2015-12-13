This is a demonstration of how two agents that have a pair of RSA keys for encryption
can negotiate a new RSA keypair (or virtual keypair) for which neither one of the keys is public.
This allows for one agent to be authorized to write, while the other is authorized to read (usually a secret key),
such that it is not possible for the reading agent to ignore the validity of a signature and use the data anyway.

This also looks like a PGP like scheme in that we can have data encrypted under a secret key paired up
with its key encrypted to somebody granted access.  But in this case, only an authority can encrypt these keys
while only users can decrypt them, which means that users cant grant access to each other, while the authority
(a program) needs to ask a user with access to the key so that the program can grant to someone else.

It is an attempt to deal with the problem of only giving a user access in the context of a program, without giving either one sufficient access on their own.
