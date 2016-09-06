#ifndef GOSTEEMRPC_SIGNING_H
#define GOSTEEMRPC_SIGNING_H

int sign_transaction(
	const unsigned char *digest,
	const unsigned char *privkey,
	unsigned char *signature,
	int *recid
);

// pubkey is expected to be 33 bytes long so that a compressed public key fits.
int verify_recoverable_signature(
	const unsigned char *digest,
	const unsigned char *signature,
	int recid,
	unsigned char *pubkey
);

#endif
