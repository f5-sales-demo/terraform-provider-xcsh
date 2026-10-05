---
page_title: "Release history"
subcategory: ""
description: "Published stable provider release history."
xcsh_docs: {"aliases": ["changelog", "provider releases", "release history"], "body_bytes": 15091, "body_sha256": "sha256:3190d15c98df499f8a946c27404c830d387b3adc8815ec57eed9d6508fde79f0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved"}, "collection_id": "xcsh-docs:guides:release-history:collection", "completeness": "complete", "id": "xcsh-docs:guides:release-history:overview", "parent_id": "xcsh-docs:provider:xcsh:navigation", "path": "documentation/guides/release-history/index.md", "product": "distributed-cloud", "provider_name": "release-history", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "guides", "registry_path": "docs/guides/release-history.md", "relationships": [], "retrieval_version": 1, "role": "overview", "schema_path": [], "schema_version": 1, "sections": [], "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Published stable provider release history.", "tasks": []}
---

# Release history

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v14.0.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v14.0.0) - 2026-10-05

Published from [`961c1db17589`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/961c1db17589302b79d8e8fa11d0e4f5a5b9beab).

### Breaking changes

API v11.0.0 retires these ten Terraform types:

- resource `xcsh_aws_vpc_site`
- resource `xcsh_azure_vnet_site`
- resource `xcsh_cloud_connect`
- resource `xcsh_gcp_vpc_site`
- resource `xcsh_securemesh_site`
- data source `xcsh_aws_vpc_site`
- data source `xcsh_azure_vnet_site`
- data source `xcsh_cloud_connect`
- data source `xcsh_gcp_vpc_site`
- data source `xcsh_securemesh_site`

Review existing configurations and state before upgrading. SMSv2 and AWS TGW are separate interfaces; this release does not establish a drop-in replacement for the retired types.

### Changes

- Aligned the provider with API v11.0.0 ([#2426](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2426)).

## [v13.1.1](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v13.1.1) - 2026-10-04

Published from [`b35faa68ef9f`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/b35faa68ef9f230fa8299444aaceddab43150715).

### Changes

- Reconciled governed repository files ([#2424](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2424)).
- Published the pinned specification delivery receipt ([#2420](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2420)).

## [v13.1.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v13.1.0) - 2026-10-04

Published from [`6e78927a063c`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/6e78927a063ce2cded2dc9967f21e78761d5b5c1).

### Changes

- Updated the F5 Distributed Cloud OpenAPI specifications ([#2416](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2416)).

## [v13.0.3](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v13.0.3) - 2026-10-03

Published from [`b1854d740269`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/b1854d740269b420e238aed9e8a556d0bb084164).

### Changes

- Cleared computed default OneOf markers for explicit selections ([#2412](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2412)).

## [v13.0.2](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v13.0.2) - 2026-10-03

Published from [`6d488749cd89`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/6d488749cd8995795a91e43b9a1892866fd73f87).

### Changes

- Reconciled governed repository files ([#2410](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2410)).

## [v13.0.1](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v13.0.1) - 2026-10-03

Published from [`e441c1d9d6ff`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/e441c1d9d6ffbc1818177215870161605e9deab7).

### Changes

- Retired historical documentation staging ([#2406](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2406)).

## [v13.0.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v13.0.0) - 2026-10-03

Published from [`17e31c1e5256`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/17e31c1e5256dbcacf2af1be4aae40348f883b0e).

### Changes

- Updated the F5 Distributed Cloud OpenAPI specifications with breaking contract changes ([#2401](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2401)).

## [v12.4.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.4.0) - 2026-10-03

Published from [`c0169b220fe4`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/c0169b220fe41a707260a88378dfb4d0b339f9b3).

### Changes

- Preserved independently reviewed property retrieval intent ([#2382](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2382)).
- Updated the F5 Distributed Cloud OpenAPI specifications ([#2393](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2393)).
- Scoped lifecycle timeout aliases to exact schema paths ([#2379](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2379)).
- Restored owned terminal KVM SLI during teardown ([#2387](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2387)).

## [v12.3.2](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.3.2) - 2026-10-03

Published from [`45349ee11ebc`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/45349ee11ebcc92c3314552fd97dc184ed1decc6).

### Changes

- Published granular canonical LLM taxonomy ([#2377](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2377)).

## [v12.3.1](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.3.1) - 2026-10-02

Published from [`2f59e49f839f`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/2f59e49f839f5da493bc716669c427aa0f16445a).

### Changes

- Scoped login success aliases to outcome schema paths ([#2374](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2374)).

## [v12.3.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.3.0) - 2026-10-02

Published from [`4bec1295c529`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/4bec1295c529c263f4fff604f8b28904795c8559).

### Changes

- Added management of regional availability for allocated public IPs ([#2368](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2368)).
- Scoped retrieval aliases and verified type choices ([#2367](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2367)).

## [v12.2.1](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.2.1) - 2026-10-02

Published from [`8d2c4d1b067b`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/8d2c4d1b067b3dadb5887cdf961670d8dae2a2b1).

### Changes

- Corrected Registry provider requirement syntax in lifecycle tests ([#2364](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2364)).

## [v12.2.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.2.0) - 2026-10-02

Published from [`b53aad4062cc`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/b53aad4062cce8c72c387720a614043232d129b9).

### Changes

- Enriched progressive Terraform retrieval metadata ([#2349](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2349)).
- Replaced immutable HTTP load balancer type selections ([#2345](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2345)).
- Preserved patch contract validation for API v9.0.1 delivery ([#2351](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2351)).

## [v12.1.2](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.1.2) - 2026-10-02

Published from [`38b94aa2e717`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/38b94aa2e7178544f29f4df416772b744ab47c9b).

### Changes

- Retired the bespoke documentation renderer after Astro Pages acceptance ([#2343](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2343)).

## [v12.1.1](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.1.1) - 2026-10-01

Published from [`c167638c2b56`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/c167638c2b56cabb1e46905609ebf9b21e4c9ba8).

### Changes

- Reconciled governed repository files ([#2337](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2337)).

## [v12.1.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.1.0) - 2026-10-01

Published from [`222209fcc6fa`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/222209fcc6fa40d43c666a38d8183ea19ec406f9).

### Changes

- Preserved the canonical documentation corpus with a compact Registry projection ([#2335](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2335)).

## [v12.0.7](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.7) - 2026-10-01

Published from [`3ae9d3a746c3`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/3ae9d3a746c330b2d8d968169c7797a3a2eb2675).

### Changes

- Published canonical Terraform documentation snapshots ([#2334](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2334)).

## [v12.0.6](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.6) - 2026-10-01

Published from [`0fd1299c89d9`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/0fd1299c89d9a7277f7550823a521e257dd9a480).

### Changes

- Corrected the namespace import contract and lookup scope ([#2331](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2331)).

## [v12.0.5](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.5) - 2026-10-01

Published from [`a2ce72095a5e`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/a2ce72095a5efd4007db3fbcaf2ad675287b03e0).

### Changes

- Published immutable Terraform documentation snapshots ([#2329](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2329)).

## [v12.0.4](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.4) - 2026-10-01

Published from [`1c4de183589a`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/1c4de183589a78ae6e52d35fca3da656c7749743).

### Changes

- Generated complete, progressively linked provider documentation ([#2288](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2288)).
- Normalized canonical collections for Registry and Pages publication ([#2292](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2292)).
- Restored complete tracked documentation and guarded manifest-owned staging ([#2323](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2323)).

## [v12.0.3](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.3) - 2026-09-30

Published from [`4bfb31e822ce`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/4bfb31e822ce43744de68a8646fbad2688fa5d60).

### Changes

- Preserved realized KVM nodes during site metadata updates ([#2282](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2282)).

## [v12.0.2](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.2) - 2026-09-30

Published from [`50775007b581`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/50775007b58176ba20c2aa1b2305d4bde265baf3).

### Changes

- Reconciled governed repository files ([#2284](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2284)).

## [v12.0.1](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.1) - 2026-09-30

Published from [`4cce1ad40637`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/4cce1ad40637d77c1c25a4becd2d7225a950f0a4).

### Changes

- Added non-publishing release transactions and retained benchmark evidence ([#2257](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2257)).
- Authorized signed tags through workflow history ([#2273](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2273)).

## [v12.0.0](https://github.com/f5-sales-demo/terraform-provider-xcsh/releases/tag/v12.0.0) - 2026-09-30

Published from [`13779f72723c`](https://github.com/f5-sales-demo/terraform-provider-xcsh/commit/13779f72723cc88d7fd0e9cfeb2f7c09f4f61daf).

### Changes

- Updated to the API v9 contract and expanded the provider surface ([#2247](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2247)).
- Moved provider validation to the verified generation artifact and governed runner flow ([#2253](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2253)).

## Historical notes

### Restored - v11.4.0

- Restored every resource, data source, and action type registered by v9.5.2 while retaining the v11 SMSv2 lifecycle and offline network allowlist data sources.
- Added an installed-schema compatibility gate and a representative CSD configuration test to prevent release-surface contractions.
- Documented the v10.0.0 through v11.3.0 broad-provider regression. Consumers upgrading from v9 should review a saved plan and investigate every managed-resource change before applying; no state surgery is expected or supported.
- Kept SMSv2 upgrade operations system-scoped and non-forced. Their v11 action inputs (`site`, `software_version`, and `os_version`) replace the v9 generic `namespace`, `name`, `version`, and `force` fields.

### Breaking Changes - v3.0.0 Clean Break Release

This is a clean-break prerelease that requires recreating all Terraform-managed resources. This version uses F5 Distributed Cloud API v2 specifications and removes all backwards compatibility with earlier versions.

#### What Changed

- **API Version**: Migrated from F5XC API v1 to API v2 specifications
- **Resource Count**: Reduced from 146 to 98 resources (removed resources without v2 API specs)
- **State Management**: Removed all state upgrade infrastructure
  - No automatic migration from previous versions
  - No schema versioning
  - No private state metadata for drift detection
  - `terraform import` still supported for adopting existing F5XC resources

#### Migration Required

Since this is a prerelease project, users must:

1. **Destroy all existing resources**: Run `terraform destroy` with the previous provider version
2. **Upgrade provider**: Update to v3.0.0 in your Terraform configuration
3. **Reinitialize**: Run `terraform init -upgrade`
4. **Recreate resources**: Run `terraform apply` to create fresh resources

**Note**: This will cause downtime. Plan accordingly.

#### Alternative: Import Existing Resources

If you have existing F5 Distributed Cloud resources (not managed by Terraform), you can import them:

```bash
terraform import xcsh_namespace.example my-namespace
terraform import xcsh_http_loadbalancer.example namespace/loadbalancer-name
```

See the provider documentation for resource-specific import formats.

### Added

- 98 resources based on F5 Distributed Cloud API v2 specifications
- Provider-defined functions:
  - `provider::xcsh::blindfold` - Encrypt plaintext with F5XC blindfold encryption
  - `provider::xcsh::blindfold_file` - Encrypt file contents with F5XC blindfold encryption

### Improved

- Cleaner codebase with ~1,000+ lines of migration code removed
- Simpler maintenance without backwards compatibility infrastructure
- Faster CI without state upgrade test coverage
- Consistent resource schemas based on OpenAPI v2 specifications

### Removed

- 48 resources without F5 Distributed Cloud API v2 specifications
- State upgrade framework (`internal/stateupgraders/`)
- Private state metadata package (`internal/privatestate/`)
- Resource upgrade tool (`tools/upgrade-resources.go`)
- Schema versioning constants
- UpgradeState methods from all resources
