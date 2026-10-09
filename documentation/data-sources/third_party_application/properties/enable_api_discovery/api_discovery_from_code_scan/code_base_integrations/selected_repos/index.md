---
page_title: "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos"
subcategory: ""
description: "Select which API repositories represent the LB applications."
xcsh_docs: {"aliases": ["enable api discovery api discovery from code scan code base integrations selected repos"], "body_bytes": 1544, "body_sha256": "sha256:3f529e9c178adf799e03ce737a3e009323fcf241040ad550da3def91e2c758ef", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "path": "documentation/data-sources/third_party_application/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2003201213002303-3110330110131020-1310203223000210-1022330310330223-3311222013102021-3230133230311332-1332231000031110-0302030303032201", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api discovery from code scan code base integrations selected repos api code repo"], "anchor": "schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo", "description": "Code repository which contain API endpoints.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos", "api_code_repo"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Select which API repositories represent the LB applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/)
- [enable_api_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_discovery_from_code_scan/)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="section"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

## Direct properties

<a id="schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo"></a>

### api_code_repo property

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.
