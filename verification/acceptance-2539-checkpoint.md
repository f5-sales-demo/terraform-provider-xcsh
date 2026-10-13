# Issue 2539 verification checkpoint

The acceptance repair is in progress. Keep issue 2539 open.

The original run is 38223043899 at 0b26d94a220f21580a4cc579327e178f01280d84.
Its real API report archive SHA-256 is
5999c68ecf47240161829adf0cc24debfa09b266a124a1914ef158885ce7e831.
The baseline is 657 tests: 309 passed, 112 failed, 236 skipped.

Original-runner reproduction 38330290832 confirmed the BGP unsupported argument
and missing service policy rule action with Go 1.25.13 and repository secrets.

Focused live run 38333807450 passed two supported tests, skipped one on an
operation-specific namespace POST 403, and reported no failures. Its ownership
receipt confirmed two resources absent and a second empty verification.

Earlier focused run 38332478097 exposed the system_metadata.uid response shape.
Recovery run 38333467770 reconstructed the original ownership journal only after
matching its SHA-256, then confirmed both owned resources absent in two passes.

The local installed v15.7.2 schema audit passed all 377 helper configurations.
Reporter, capability, ownership, and workflow support regressions passed.
Root block-conflict generation requires production regeneration and release
qualification before delivery. Full mock/live comparison, all prerequisite
receipts, PR completion, CI, merge, release, and protected-main acceptance remain.
