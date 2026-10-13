# Issue 2539 verification checkpoint

The acceptance repair is in progress. Keep issue 2539 open and PR 2544 in draft.

The preserved baseline is run 38223043899 at
0b26d94a220f21580a4cc579327e178f01280d84: 657 tests, 309 passed, 112 failed,
236 skipped. Its archive SHA-256 is
5999c68ecf47240161829adf0cc24debfa09b266a124a1914ef158885ce7e831.

The ownership journal, generated read/delete handlers, exact operation inventory,
fixture validation, and strict reporter are implemented. Cleanup executes in the
real-test job and uses typed absence in two passes. Unknown output and incomplete
streams remain failures. Test-issued deletes now require exact journal ownership
and matching UID readback; loaded journals reject raw UIDs, unconfirmed created
identities, and unavailable verifiers.

Reporter regressions reproduced missing package starts and missing Markdown on
incomplete mock streams; the repairs passed. Focused dispatch arguments are
transported through environment variables. Datasource fixture namespaces now use
the provider's verified spec constraint resolver. Appliance references use exact
read-only API readback and remain private.

Measured live replays:

| Run | Passed | Failed | Skipped | Owned objects absent twice |
| --- | ---: | ---: | ---: | ---: |
| 38333807450 | 2 | 0 | 1 | 2 |
| 38334365131 | 276 | 90 | 291 | 307 |
| 38347540472 | 2 | 1 | 0 | 5 |
| 38349354346 | 5 | 2 | 1 | 10 |
| 38341637137 | 5 | 1 | 0 | 8 |

WAF import verification, HTTP load balancer lifecycle, cluster groups, protocol
policers, and workload flavors passed focused live tests. All eight API-definition lifecycle tests passed in run 38356340525 after
compact identities satisfied the measured 20-character API name limit. The latest mesh-group replay returned HTTP 400 for a missing
input; its verified virtual-site selector now skips typed stale inventory entries. Three
mesh lifecycle cases still fail HTTP 400 for a required input. That run confirmed
20 owned objects absent twice. BGP skipped on an operation-specific POST 403.

Older full comparison 38337641132 timed out after two hours with 187 unfinished
tests; its incomplete artifact is preserved. The older local mock completed with
664 tests: 610 passed, 17 failed, 37 skipped. All 17 failures have subsequent focused
repairs; this is not final acceptance.

Go 1.25.13 support regressions, installed-schema fixture audits, focused mocks,
namespace/schema/codegen tests, and the resumed broad lint passed. Full serial
mock/live acceptance on final source, evidence for every skip, final CI, release
and immutable artifact installation, and protected-main acceptance remain required.
Production generator changes require a new qualified provider release.

Docs-control PR 2489 and manifest PR 2490 merged. Provider reconciliation PR 2548
pairs the canonical policy with the focused dispatch declaration; its exact trigger
and mutation tests pass under race detection. Full CI and merge remain pending.
The resumed complete serial mock and full live run 38357531909 are active.
