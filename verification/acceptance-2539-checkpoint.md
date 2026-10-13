# Issue 2539 verification checkpoint

Acceptance repair is in progress. Keep issue 2539 open and PR 2544 in draft until the full suite, cleanup, protected-main, and artifact gates pass.

The preserved baseline is run 38223043899 at source 0b26d94a220f21580a4cc579327e178f01280d84: 657 tests, 309 passed, 112 failed, and 236 skipped. Its archive SHA-256 is 5999c68ecf47240161829adf0cc24debfa09b266a124a1914ef158885ce7e831.

The implementation includes a private ownership journal, generated CRUD read/delete handlers, operation-specific prerequisite inventory, positive fixture validation, and strict sanitized reporting. Cleanup runs in the real-test workspace, deletes exact journal-owned identities in dependency order, and verifies absence twice. Unknown diagnostics, incomplete streams, missing verifiers, uncertain writes, and incomplete cleanup fail closed.

The completed serial mock checkpoint contains 664 tests: 627 passed, zero failed, and 37 skipped, with a complete stream. This checkpoint predates later fixture and prerequisite changes and does not qualify final source.

The broad live checkpoint at run 38357531909 contains 659 tests: 392 passed, 27 failed, and 240 skipped. Two passing qualification diagnostics were extra suite members, so actual acceptance coverage was 390 passes, an increase of 81 against the preserved baseline. The exact-inventory gate correctly rejected that run. Its 417 owned objects were confirmed absent twice. Final acceptance must contain exactly the 657 inventoried acceptance tests.

Subsequent focused repairs passed API-definition lifecycles, certificate/chain cases, WAF import, route, UDP, policy-view, allocator, ACL, service-policy-rule, workload-flavor, CDN, and external-connector cases. New receipts in acceptance-2539-latest-replays.json preserve the measured focused results and digests, including failures.

Disposable crawler, CRL, and registry probes returned typed HTTP 429 responses explicitly stating that the object kind exhausted limits. Forwarding-class has no entry in the typed quota inventory, and global receiver's verified count equals its configured maximum. These outcomes provide specific prerequisite evidence without inferring why collection counts and object-limit enforcement differ. Generic rate limiting, validation failures, and HTTP 500 responses remain failures.

Run 38384379873 confirmed that the segment reference emitted by the provider had a present 20-character name, valid characters, and namespace system. The API still rejected the reference with HTTP 400. Mesh creation still returned HTTP 500. All seven journal entries were confirmed absent twice. This diagnostic branch selected five tests but retained the full expected inventory; completeness correctly failed and the run is not suite qualification.

Fresh regressions reproduced acceptance of empty or failed mesh selector expansions. Mesh discovery now uses the read-only selectees endpoint and exact selected-site UID readback; empty successful expansion can establish absence, while malformed responses, authorization failures, server errors, and stale selected sites fail. Focused tests, race checks, support/reporter/comparison tests, and lint passed.

A separate regression reproduced missing BFD, fallback, mesh mode, and virtual-site reference in the generated mesh datasource fixture. Its generator now emits those fields and requires the verified read-only reference. Regeneration, the positive installed-schema audit, and focused mock passed; provider lint reported zero issues.

Canonical focused-dispatch policy delivery is complete through docs-control PRs 2489/2490 and provider reconciliation PR 2548. The canonical policy blob is 27a3e0a8a9715cec313d7c7c82a507750ec135fc.

The authorized additional tenant enabled crawler, CRL, and registry live passes. Compact site and segment identities repaired the derived segment VRF label failure; segment VRF and global log receiver passed together with complete reporting. Initial namespace cleanup failures are preserved, and recovered receipts confirmed every entry absent twice. Namespace cleanup now uses an explicit command handler after exact UID and empty referring-object verification.

Read-only existing mesh inspection and disposable probes measured data-plane mesh with regional fallback as the supported fixture. All 14 mesh lifecycle/datasource tests passed live with zero skips or failures; all 16 owned entries were confirmed absent twice. New acceptance prerequisites clone verified selector expressions into a journal-owned system selector and confirm its resolved membership through read-only appliance checks.

The resumed serial mock at source 36c8e6f206229a843c3ea898983f521a7d43d0d2 completed with 664 tests: 627 passed, zero failed, and 37 skipped. Its stream is complete. A fresh frozen qualification checkout at d4f8bff2351 includes the measured repairs and owned selector provisioning. Final-source mock/live comparison remains required.

Outstanding work: complete final-source serial mock/live comparison with exact accounting and evidence for every skip; finish CI and merge only after gates pass; run protected-main acceptance and cleanup; release and qualify the immutable provider artifact because earlier production generator behavior changed. Existing task and qualification worktrees remain preserved while their runs or evidence are needed.
