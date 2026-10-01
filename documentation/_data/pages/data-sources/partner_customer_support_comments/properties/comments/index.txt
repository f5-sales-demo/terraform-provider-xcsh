---
page_title: "comments"
subcategory: ""
description: "comments for xcsh_partner_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 2457, "body_sha256": "sha256:d9665134250b752c528bde8b85faa23f20e2a581d5118c53466a798e86d01023", "child_ids": ["xcsh-docs:data-sources:partner_customer_support_comments:properties:comments:attachments_info"], "collection_id": "xcsh-docs:data-sources:partner_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:partner_customer_support_comments:properties:comments", "parent_id": "xcsh-docs:data-sources:partner_customer_support_comments:reference", "path": "documentation/data-sources/partner_customer_support_comments/properties/comments/index.md", "provider_name": "partner_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["comments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/partner_customer_support_comments/properties/comments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "comments for xcsh_partner_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments

Breadcrumbs:

- [xcsh_partner_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/)
- comments

<a id="section"></a>

Type: `"list"`. Computed.

List of comments on the customer support ticket.

## Direct properties

<a id="schema-comments--attachment_ids"></a>

### attachment_ids property

Type: `["list", "string"]`. Computed.

Third party ID of any attachment related to this ticket comment.

- [attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/): complete subsection reference.

<a id="schema-comments--author_email"></a>

### author_email property

Type: `"string"`. Computed.

Email. Email of the author of the comment.

<a id="schema-comments--author_name"></a>

### author_name property

Type: `"string"`. Computed.

Author. Author of the comment (as a name)

<a id="schema-comments--comment_id"></a>

### comment_id property

Type: `"string"`. Computed.

ID assigned to this comment by support provider.

<a id="schema-comments--created_at"></a>

### created_at property

Type: `"string"`. Computed.

At. Comment creation time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="schema-comments--html"></a>

### html property

Type: `"string"`. Computed.

Comment. Comment body as HTML.

<a id="schema-comments--plain_text"></a>

### plain_text property

Type: `"string"`. Computed.

Comment. Comment body as plain text.

## Next pages

- [comments.attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/comments/attachments_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/)
- [xcsh_partner_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/)
