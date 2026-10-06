Resource lifecycle delivery batch
=================================

Tracks #2445 and #2330. Source base: 963f5b60542.
API producer: immutable v12.0.0 at a9be0360815e3fd3fa08ded845e03c0bd4afd6ec.

- [x] Fresh Ubuntu worktree and current issue/source intake.
- [x] Close verified intake issues #2457, #2081, #2332, #2311, #2301, #1992.
- [x] Reproduce live CSD GET 501 and exact-root collection projection without mutation.
- [x] Implement typed status and authoritative client reconciliation.
- [x] Focused client and resource refresh/import/deletion regressions pass.
- [x] Generator persistence regressions pass; root lookup implemented.
- [x] Real Terraform WAF saved-plan transitions, payload exclusivity and concurrency pass.
- [ ] Full source, race, lint, privacy and deterministic generation gates.
- [ ] Fresh reviewed CSD and WAF plans, live acceptance and convergence.
- [ ] Source PR, regeneration, immutable release, Pages and Registry acceptance.
- [ ] Installed successor acceptance and issue closures.
- [ ] Assess remaining automated/governance issues against concrete evidence.

The original protected_domain.List RPC describes the set in one namespace,
repeated report_fields, items[].get_spec and collection errors. It has no
pagination request or continuation response. The inferred 256-item enrichment
bound does not define a truncation contract. Unknown envelope fields, missing
items, collection errors, incomplete roots, duplicate roots and conflicting
scope/identity fail closed. An explicit empty collection is authoritative under
this complete RPC contract. Fixtures exercise counts 255, 256 and 257.

The live backend returns the configured root but can omit name, namespace, UID
and metadata. Terraform identity remains configured ownership; no backend ID is
invented. Blank metadata cannot establish metadata drift or console visibility.
The lookup reuses its existing protected_domain field as an optional root input.

CSD acceptance uses its existing encrypted S3 backend and lockfile. WAF uses the
remediation-553 workspace and existing private application state. Historical
plans are evidence only. Generate, review and hash each plan before applying.
Preserve unrelated cloud objects, namespace, alert settings and traffic workers.
Keep csd#1296 open until independent console/API parity passes.

Fresh CSD baseline: zero changes and zero drift. SHA-256
6835bb2c6e75fcbef52ef5f9cb5a1c44290ecf6adf51f4ba7a85cd261b393479.
Fresh WAF baseline: firewall no-op; two unrelated VM custom_data replacements
were rejected for apply. SHA-256
e3677160cc417f5f87dc172c6472de2d5962f07665cea3665464aa1eb602df49.
Ruff, MyPy and Pylint pass. Full Go race/lint and docs gates remain active.

CSD live acceptance: authoritative import and two no-op apply cycles passed.
Controlled name-key DELETE returned 200 without deletion; exact root-key DELETE
removed the registration. Fresh plan proposed only xcsh_protected_domain.csd
create, saved-plan apply passed, then full no-op/zero-drift convergence passed.
The client now uses the observed root-key deletion contract and verifies absence
before reporting deletion success. Terraform name/ID ownership is unchanged.

WAF live acceptance: custom to omitted default, default to custom, custom to
explicit default, and restore custom passed. Every saved plan changed only the
firewall address in place; UID stayed unchanged and each subsequent plan was no-op.
Temporary test override was removed; application workspace is clean.

Full Go race and lint passed before the final domain-key deletion adjustment.
Final deletion regressions/client tests and staged PII pass; full final gates run.
All 330 release-surface and 76 maintained examples plus schema compatibility pass.
Biome 2.5.6 is task-private; host 2.5.8 mismatch required a docs rerun.
Console browser inventory is empty; csd#1296 remains open.


Final verification update: full Go race suite with the domain-key delete contract
passed. Documentation pipeline passed with exact Biome 2.5.6: 330 collections,
16,366 pages, 34,603 exact property destinations; 1,426 Registry documents,
maximum 426,017 bytes, zero oversized exceptions; coverage, links/fragments,
compatibility and deterministic regeneration passed. Corpus Markdown lint passed across 17,792 generated files. Final Go lint passed with zero issues. Final candidate SHA-256:
f969f25de40e951b21cfa1e96b22c7b95e9d45e5362bd0ac4a6809c4dbcfbb44.
Final candidate refresh yields CSD full-stack and WAF firewall-only no-op plans.

All required local source gates and scoped live lifecycle checks passed. Publication
and installed immutable-release acceptance remain required before issue closure.
