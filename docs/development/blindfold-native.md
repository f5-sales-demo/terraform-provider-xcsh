# Native Blindfold certificate interface

The optional `blindfold` object is available beside `certificate_url` and `private_key` at every
recognized certificate node. Native mode conflicts with supplied certificate locations and key
blocks. Existing supplied-location configurations remain supported.

Use `certificate_file` with `private_key_file`, or `pkcs12_file`. Value input uses public
`certificate_pem` and sensitive write-only `private_key_wo`, or sensitive write-only base64
`pkcs12_wo`. Protected input uses `passphrase_env` or sensitive write-only `passphrase_wo`.
Write-only keys and bundles require `material_version`; Terraform 1.11 or later is required.
Inline nodes require a stable unique `id`; named certificates default to `default`. Policy defaults
to `shared/ves-io-allow-volterra`.

Computed values include `fingerprint`, `expires_at`, `context_digest`, `chain_identity`,
`spki_identity`, `algorithm`, `prepared_identity`, and sensitive `encrypted_location`.
Normalization happens in Go. Encryption happens during apply. There are no runtime downloads,
xcsh calls, or external OpenSSL requirements. PKCS#8 and modern P12 dependencies are module-pinned.

Remote reads use complete replace-form configuration and retain the optimistic concurrency token.
Unchanged public identity, context and material version reuse ciphertext; remote key drift repairs
the retained ciphertext. Saved-plan material or context changes require replanning. In-place
RSA/EC changes require a new certificate/reference cutover. Import does not recover private inputs.
Provenance annotations are reconciliation metadata and contain only public identities and ciphertext
hashes. Sequential tool handoff is supported; concurrent ownership is unsupported.

Qualification is tracked in the acceptance receipt. A schema surface or unit test does not establish
live TLS qualification for that surface. Public Terraform guide publication follows installed/live
acceptance and is outside this iteration.
