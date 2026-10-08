---
page_title: "filters.global_filters"
subcategory: ""
description: "Global Filters. List of global filters."
xcsh_docs: {"aliases": ["filters global filters"], "body_bytes": 6622, "body_sha256": "sha256:8471f246ded822edfa74ca1e0bf94233be3315d81f081fe412a44aad2ca4b3fd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters:global_filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters", "path": "documentation/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1123331132203332-2000203310223103-0011123103333212-3321110311302130-3122310023020331-2022200312111212-0100302221232323-0132123001120231", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filters", "global_filters"], "schema_version": 1, "sections": [{"aliases": ["filters global filters key"], "anchor": "schema-filters--global_filters--key", "description": "Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible values are `TIMESTAMP`, `USERNAME`, `CLIENT_TOKEN`, `IP_ADDRESS`, `ASN`, `AS_ORGANIZATION`, `COUNTRY`, `METHOD`, `HOST`, `PATH`, `URL`, `REFERER`, `TRAFFIC_CHANNEL`, `IS_ATTACK`, `BOT_REASON`, `TRAFFIC_TYPE`, `THREAT_TYPE`,", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters:global_filters", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ABSOLUTE", "ACTION_TAKEN", "AGENT", "APPLICATION_NAME", "ASN", "AS_ORGANIZATION", "BOT_COOKIE", "BOT_ENDPOINT_POLICY", "BOT_REASON", "BROWSER_FINGERPRINT", "CLIENT_TOKEN", "COOKIE_AGE", "COUNTRY", "DEVICE_ID", "ENDPOINT_LABEL", "ENDPOINT_NAME", "ENDPOINT_POLICY", "FLOW", "FLOW_CATEGORY", "FLOW_LABEL", "HEADER_FINGERPRINT", "HOST", "IP_ADDRESS", "IS_ATTACK", "KNOWN_BOT_CATEGORY", "KNOWN_BOT_CATEGORY_TYPE", "KNOWN_BOT_MITIGATION", "KNOWN_BOT_NAME", "KNOWN_BOT_PROVIDER", "METHOD", "MOBILE_TRANSACTION_INSIGHT", "PATH", "PERCENTAGE", "PROTECTED_APPLICATION", "REFERER", "RESPONSE_CODE", "SDK_VERSION", "SERVER_RESPONSE_CODE", "THREAT_TYPE", "TIMESTAMP", "TRAFFIC_CHANNEL", "TRAFFIC_TYPE", "TRANSACTION_RESULT", "TREND", "TRIGGERED_RULE", "URL", "USERNAME", "USER_AGENT", "USER_AGENT_FAMILY", "USER_AGENT_OS_FAMILY", "USER_FINGERPRINT", "WEB_TRANSACTION_INSIGHT"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "global_filters", "key"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters global filters op"], "anchor": "schema-filters--global_filters--op", "description": "Operator for query filter - IN: Filter Operator Specifies that query result includes filter values - NOT_IN: Filter Operator Specifies that query result excludes filter values - MATCHES_REGEX: Filter Operator Specifies that query result matches filter regex - DOES_NOT_MATCH_REGEX: Filter.. Possible values are `IN`,", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters:global_filters", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DOES_NOT_INCLUDE", "DOES_NOT_MATCH_REGEX", "ENDS_WITH", "IN", "INCLUDES", "MATCHES_REGEX", "NOT_IN", "STARTS_WITH"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "global_filters", "op"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters global filters values"], "anchor": "schema-filters--global_filters--values", "description": "Values. An unordered list of filter strings.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters:global_filters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "global_filters", "values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Global Filters. List of global filters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filters.global_filters

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/)
- filters.global_filters

<a id="section"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

## Direct properties

<a id="schema-filters--global_filters--key"></a>

### key property

Type: `"string"`. Optional.

\[Enum:
TIMESTAMP|USERNAME|CLIENT\_TOKEN|IP\_ADDRESS|ASN|AS\_ORGANIZATION|COUNTRY|METHOD|HOST|PATH|URL|REFERER|TRAFFIC\_CHANNEL|IS\_ATTACK|BOT\_REASON|TRAFFIC\_TYPE|THREAT\_TYPE|SDK\_VERSION|ACTION\_TAKEN|COOKIE\_AGE|BOT\_COOKIE|USER\_AGENT|USER\_AGENT\_OS\_FAMILY|USER\_AGENT\_FAMILY|BROWSER\_FINGERPRINT|USER\_FINGERPRINT|HEADER\_FINGERPRINT|DEVICE\_ID|FLOW|AGENT|APPLICATION\_NAME|PROTECTED\_APPLICATION|RESPONSE\_CODE|SERVER\_RESPONSE\_CODE|TRANSACTION\_RESULT|MOBILE\_TRANSACTION\_INSIGHT|WEB\_TRANSACTION\_INSIGHT|TRIGGERED\_RULE|FLOW\_CATEGORY|FLOW\_LABEL|ENDPOINT\_NAME|ENDPOINT\_LABEL|BOT\_ENDPOINT\_POLICY|KNOWN\_BOT\_NAME|KNOWN\_BOT\_CATEGORY|KNOWN\_BOT\_PROVIDER|KNOWN\_BOT\_CATEGORY\_TYPE|KNOWN\_BOT\_MITIGATION|ABSOLUTE|PERCENTAGE|TREND|ENDPOINT\_POLICY\]
Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible
values are \`TIMESTAMP\`, \`USERNAME\`, \`CLIENT\_TOKEN\`, \`IP\_ADDRESS\`, \`ASN\`,
\`AS\_ORGANIZATION\`, \`COUNTRY\`, \`METHOD\`, \`HOST\`, \`PATH\`, \`URL\`, \`REFERER\`,
\`TRAFFIC\_CHANNEL\`, \`IS\_ATTACK\`, \`BOT\_REASON\`, \`TRAFFIC\_TYPE\`, \`THREAT\_TYPE\`,
\`SDK\_VERSION\`, \`ACTION\_TAKEN\`, \`COOKIE\_AGE\`, \`BOT\_COOKIE\`, \`USER\_AGENT\`,
\`USER\_AGENT\_OS\_FAMILY\`, \`USER\_AGENT\_FAMILY\`, \`BROWSER\_FINGERPRINT\`,
\`USER\_FINGERPRINT\`, \`HEADER\_FINGERPRINT\`, \`DEVICE\_ID\`, \`FLOW\`, \`AGENT\`,
\`APPLICATION\_NAME\`, \`PROTECTED\_APPLICATION\`, \`RESPONSE\_CODE\`, \`SERVER\_RESPONSE\_CODE\`,
\`TRANSACTION\_RESULT\`, \`MOBILE\_TRANSACTION\_INSIGHT\`, \`WEB\_TRANSACTION\_INSIGHT\`,
\`TRIGGERED\_RULE\`, \`FLOW\_CATEGORY\`, \`FLOW\_LABEL\`, \`ENDPOINT\_NAME\`, \`ENDPOINT\_LABEL\`,
\`BOT\_ENDPOINT\_POLICY\`, \`KNOWN\_BOT\_NAME\`, \`KNOWN\_BOT\_CATEGORY\`, \`KNOWN\_BOT\_PROVIDER\`,
\`KNOWN\_BOT\_CATEGORY\_TYPE\`, \`KNOWN\_BOT\_MITIGATION\`, \`ABSOLUTE\`, \`PERCENTAGE\`, \`TREND\`,
\`ENDPOINT\_POLICY\`. Defaults to \`TIMESTAMP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ABSOLUTE","ACTION_TAKEN","AGENT","APPLICATION_NAME","ASN","AS_ORGANIZATION","BOT_COOKIE","BOT_ENDPOINT_POLICY","BOT_REASON","BROWSER_FINGERPRINT","CLIENT_TOKEN","COOKIE_AGE","COUNTRY","DEVICE_ID","ENDPOINT_LABEL","ENDPOINT_NAME","ENDPOINT_POLICY","FLOW","FLOW_CATEGORY","FLOW_LABEL","HEADER_FINGERPRINT","HOST","IP_ADDRESS","IS_ATTACK","KNOWN_BOT_CATEGORY","KNOWN_BOT_CATEGORY_TYPE","KNOWN_BOT_MITIGATION","KNOWN_BOT_NAME","KNOWN_BOT_PROVIDER","METHOD","MOBILE_TRANSACTION_INSIGHT","PATH","PERCENTAGE","PROTECTED_APPLICATION","REFERER","RESPONSE_CODE","SDK_VERSION","SERVER_RESPONSE_CODE","THREAT_TYPE","TIMESTAMP","TRAFFIC_CHANNEL","TRAFFIC_TYPE","TRANSACTION_RESULT","TREND","TRIGGERED_RULE","URL","USERNAME","USER_AGENT","USER_AGENT_FAMILY","USER_AGENT_OS_FAMILY","USER_FINGERPRINT","WEB_TRANSACTION_INSIGHT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TIMESTAMP",
    "USERNAME",
    "CLIENT_TOKEN",
    "IP_ADDRESS",
    "ASN",
    "AS_ORGANIZATION",
    "COUNTRY",
    "METHOD",
    "HOST",
    "PATH",
    "URL",
    "REFERER",
    "TRAFFIC_CHANNEL",
    "IS_ATTACK",
    "BOT_REASON",
    "TRAFFIC_TYPE",
    "THREAT_TYPE",
    "SDK_VERSION",
    "ACTION_TAKEN",
    "COOKIE_AGE",
    "BOT_COOKIE",
    "USER_AGENT",
    "USER_AGENT_OS_FAMILY",
    "USER_AGENT_FAMILY",
    "BROWSER_FINGERPRINT",
    "USER_FINGERPRINT",
    "HEADER_FINGERPRINT",
    "DEVICE_ID",
    "FLOW",
    "AGENT",
    "APPLICATION_NAME",
    "PROTECTED_APPLICATION",
    "RESPONSE_CODE",
    "SERVER_RESPONSE_CODE",
    "TRANSACTION_RESULT",
    "MOBILE_TRANSACTION_INSIGHT",
    "WEB_TRANSACTION_INSIGHT",
    "TRIGGERED_RULE",
    "FLOW_CATEGORY",
    "FLOW_LABEL",
    "ENDPOINT_NAME",
    "ENDPOINT_LABEL",
    "BOT_ENDPOINT_POLICY",
    "KNOWN_BOT_NAME",
    "KNOWN_BOT_CATEGORY",
    "KNOWN_BOT_PROVIDER",
    "KNOWN_BOT_CATEGORY_TYPE",
    "KNOWN_BOT_MITIGATION",
    "ABSOLUTE",
    "PERCENTAGE",
    "TREND",
    "ENDPOINT_POLICY"),
}
```

<a id="schema-filters--global_filters--op"></a>

### op property

Type: `"string"`. Optional.

\[Enum:
IN|NOT\_IN|MATCHES\_REGEX|DOES\_NOT\_MATCH\_REGEX|INCLUDES|DOES\_NOT\_INCLUDE|STARTS\_WITH|ENDS\_WITH\]
Operator for query filter - IN: Filter Operator Specifies that query result includes filter values -
NOT\_IN: Filter Operator Specifies that query result excludes filter values - MATCHES\_REGEX: Filter
Operator Specifies that query result matches filter regex - DOES\_NOT\_MATCH\_REGEX: Filter..
Possible values are \`IN\`, \`NOT\_IN\`, \`MATCHES\_REGEX\`, \`DOES\_NOT\_MATCH\_REGEX\`,
\`INCLUDES\`, \`DOES\_NOT\_INCLUDE\`, \`STARTS\_WITH\`, \`ENDS\_WITH\`. Defaults to \`IN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DOES_NOT_INCLUDE","DOES_NOT_MATCH_REGEX","ENDS_WITH","IN","INCLUDES","MATCHES_REGEX","NOT_IN","STARTS_WITH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("IN",
    "NOT_IN",
    "MATCHES_REGEX",
    "DOES_NOT_MATCH_REGEX",
    "INCLUDES",
    "DOES_NOT_INCLUDE",
    "STARTS_WITH",
    "ENDS_WITH"),
}
```

<a id="schema-filters--global_filters--values"></a>

### values property

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```
