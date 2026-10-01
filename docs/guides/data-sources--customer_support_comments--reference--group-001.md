---
page_title: "xcsh_customer_support_comments reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_customer_support_comments reference."
---

# xcsh_customer_support_comments reference

<a id="canonical-12e9887903e3773a47890b40be28cd522622da1fd04215c1250c156eea0f4f47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e8fd527303dfa2ecfdc0407e9101358b334af232cb20f4b7e516f7defb40db9"></a>

## Property reference — Property reference / f640bb9e33f4 / 2

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5)
- Property reference

<a id="canonical-403fceda3677852083840f6a744902906942d5c1a2c3b1222a40880e38723516"></a>

## Direct properties — Property reference / f640bb9e33f4 / 3

<a id="canonical-5971af7c0c3dbe81db30bf2732cfe1e90422a962a3f9aa7e35c9d26076fe310b"></a>

<a id="canonical-09746a0c18e5f8712f3e955cc4815d1ee22a241dfd8c5e1481defad7f722dfd3"></a>

## all_comments_returned property — Property reference / f640bb9e33f4 / 4

Type: `"bool"`. Computed.

Indicates if all comments for the issue have been returned.

- [comments](data-sources--customer_support_comments--reference--group-001.md#canonical-260c3e6a46b7e41abb4555618a3e2f1c1fda9e189cbdfcc95b0a4a9c6f22be19): complete subsection reference.

<a id="canonical-91e9b327edde49236a9b096ba2a76418f57fc684c579069f696ac0744f58259e"></a>

<a id="canonical-9a0832dd495ce7dd89569bb2f686fd08ffc5ac4bded1c6971c55ff5ae334c3c3"></a>

## created_until_timestamp property — Property reference / f640bb9e33f4 / 5

Type: `"string"`. Optional.

Filter to retrieve comments created up to the specified timestamp.

<a id="canonical-66cdf17647b633108911338003bb9a164d7a1b1990ae18b565397c67959406a2"></a>

<a id="canonical-cb4d74513ae6dd5c00f8825d8f79abc97ea815f641aa08501e43cbda10d0fa6f"></a>

## name property — Property reference / f640bb9e33f4 / 6

Type: `"string"`. Required.

Name The name (issue ID) of the customer support ticket object.

<a id="canonical-78ea3c9ddf667b8ac4fee87690e37356ec93070c71295b7da3f3dd89a752cb5f"></a>

## All schema paths — Property reference / f640bb9e33f4 / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `all_comments_returned` | [all_comments_returned](data-sources--customer_support_comments--reference--group-001.md#canonical-5971af7c0c3dbe81db30bf2732cfe1e90422a962a3f9aa7e35c9d26076fe310b) |
| `comments` | [comments](data-sources--customer_support_comments--reference--group-001.md#canonical-3c7a0969183b1a2036b5588a5f5d21f6eb7ff2774e472bf345be61ff43eea021) |
| `comments.attachment_ids` | [comments.attachment_ids](data-sources--customer_support_comments--reference--group-001.md#canonical-d231827189a1a9b89e966658217608c419760346db1c0087008f5e78d736a0e4) |
| `comments.attachments_info` | [comments.attachments_info](data-sources--customer_support_comments--reference--group-001.md#canonical-07fab5ea44f779dc0baa7cda7a6e064fac789aabdf7d7418b84f848c73d2d711) |
| `comments.attachments_info.attachment` | [comments.attachments_info.attachment](data-sources--customer_support_comments--reference--group-001.md#canonical-da297691a5d0105dbba719a2ee23c0617c560cfa3e8454a07f1530dcdcbfa294) |
| `comments.attachments_info.content_type` | [comments.attachments_info.content_type](data-sources--customer_support_comments--reference--group-001.md#canonical-e9ff3c96a90de2343128e6e098476efab38e13570df423a8431be8423d7c5565) |
| `comments.attachments_info.filename` | [comments.attachments_info.filename](data-sources--customer_support_comments--reference--group-001.md#canonical-a88615f49e857f5747530cc252357d1f26a90bab38460ea3dde2f276b5b48326) |
| `comments.attachments_info.tp_id` | [comments.attachments_info.tp_id](data-sources--customer_support_comments--reference--group-001.md#canonical-95e09848d5ff9960f093dff4790036c49b40f8988fe1e3a09c21ca3b5de9bc6d) |
| `comments.author_email` | [comments.author_email](data-sources--customer_support_comments--reference--group-001.md#canonical-3706a10c9efb24bde7429f48494bc2fa7cd2d014daf62cf1aa081b663a2952cf) |
| `comments.author_name` | [comments.author_name](data-sources--customer_support_comments--reference--group-001.md#canonical-f1bd016f5e00c92eb4c2718c69478573387af105095a80ee07e6481c94ac2034) |
| `comments.comment_id` | [comments.comment_id](data-sources--customer_support_comments--reference--group-001.md#canonical-f05482038e9dc53b7577797106a6b615dd0123901a9b47b69bce146483c4e060) |
| `comments.created_at` | [comments.created_at](data-sources--customer_support_comments--reference--group-001.md#canonical-3031eda5d6d7c44227d1cffe82f419c351d2e35142c567a987f11ee263d9c0e3) |
| `comments.html` | [comments.html](data-sources--customer_support_comments--reference--group-001.md#canonical-a7aabf67b8760402e58af6c6bad4b7a34539e83fea1688ce6b8a7eeb87f87616) |
| `comments.plain_text` | [comments.plain_text](data-sources--customer_support_comments--reference--group-001.md#canonical-f6f3b968d77b468d16d5a390b07b73a6713c7793ca3d86e1be7946c9f5c50711) |
| `created_until_timestamp` | [created_until_timestamp](data-sources--customer_support_comments--reference--group-001.md#canonical-91e9b327edde49236a9b096ba2a76418f57fc684c579069f696ac0744f58259e) |
| `name` | [name](data-sources--customer_support_comments--reference--group-001.md#canonical-66cdf17647b633108911338003bb9a164d7a1b1990ae18b565397c67959406a2) |

<a id="canonical-2913c96071d16d29a6da53f2c1364f797e5364600fe7a44997b2a3c10e673fff"></a>

## Next pages — Property reference / f640bb9e33f4 / 8

- [comments](data-sources--customer_support_comments--reference--group-001.md#canonical-260c3e6a46b7e41abb4555618a3e2f1c1fda9e189cbdfcc95b0a4a9c6f22be19)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5)

<a id="canonical-260c3e6a46b7e41abb4555618a3e2f1c1fda9e189cbdfcc95b0a4a9c6f22be19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36d891cc1cd215dd131290b2ac48904b10eb58b3beffe0cd725e3439173833e2"></a>

## comments — comments / 0c3e6c32a37b / 2

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5)
- [Property reference](data-sources--customer_support_comments--reference--group-001.md#canonical-12e9887903e3773a47890b40be28cd522622da1fd04215c1250c156eea0f4f47)
- comments

<a id="canonical-3c7a0969183b1a2036b5588a5f5d21f6eb7ff2774e472bf345be61ff43eea021"></a>

Type: `"list"`. Computed.

List of comments on the customer support ticket.

<a id="canonical-86dce0c929ad9cf615b19fd3b600de13d8742dc69078641efeb57844042f07e0"></a>

## Direct properties — comments / 0c3e6c32a37b / 3

<a id="canonical-d231827189a1a9b89e966658217608c419760346db1c0087008f5e78d736a0e4"></a>

<a id="canonical-46ca67b46f2f4fb9c043cf5aca85d8d29cf531ea8fb6ffc2624d685829e67dff"></a>

## attachment_ids property — comments / 0c3e6c32a37b / 4

Type: `["list", "string"]`. Computed.

Third party ID of any attachment related to this ticket comment.

- [attachments_info](data-sources--customer_support_comments--reference--group-001.md#canonical-5b852e4604bfc680b986aa889b3345eca01bd273b515291bf0b8292814052821): complete subsection reference.

<a id="canonical-3706a10c9efb24bde7429f48494bc2fa7cd2d014daf62cf1aa081b663a2952cf"></a>

<a id="canonical-1e7ee26d56832239f104ab2ed24a0cc4f778cbe70a2cb64424fe4f3fea4d927e"></a>

## author_email property — comments / 0c3e6c32a37b / 5

Type: `"string"`. Computed.

Email. Email of the author of the comment.

<a id="canonical-f1bd016f5e00c92eb4c2718c69478573387af105095a80ee07e6481c94ac2034"></a>

<a id="canonical-6f1c466dde849c261c90c0ca3029b3ef4eabd8c18f97b378f2d496457805ab69"></a>

## author_name property — comments / 0c3e6c32a37b / 6

Type: `"string"`. Computed.

Author. Author of the comment (as a name)

<a id="canonical-f05482038e9dc53b7577797106a6b615dd0123901a9b47b69bce146483c4e060"></a>

<a id="canonical-1e63658e968271b182981dc0a83e20ac1d65ed1fb9a652fe8107f92c5e4f4a17"></a>

## comment_id property — comments / 0c3e6c32a37b / 7

Type: `"string"`. Computed.

ID assigned to this comment by support provider.

<a id="canonical-3031eda5d6d7c44227d1cffe82f419c351d2e35142c567a987f11ee263d9c0e3"></a>

<a id="canonical-c4eeec26d6176d6f5fcee5fd07c4cbfb87419f7b77949d366ad7a50dda0914f0"></a>

## created_at property — comments / 0c3e6c32a37b / 8

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

<a id="canonical-a7aabf67b8760402e58af6c6bad4b7a34539e83fea1688ce6b8a7eeb87f87616"></a>

<a id="canonical-02d94b00167d3c13d21fcd7f0310f74676b8d21ff7a5cbf4802bd4d90bbb3624"></a>

## html property — comments / 0c3e6c32a37b / 9

Type: `"string"`. Computed.

Comment. Comment body as HTML.

<a id="canonical-f6f3b968d77b468d16d5a390b07b73a6713c7793ca3d86e1be7946c9f5c50711"></a>

<a id="canonical-9b956a552e4105393810877bd619879b5000ec70795bd076f8e918607791324c"></a>

## plain_text property — comments / 0c3e6c32a37b / 10

Type: `"string"`. Computed.

Comment. Comment body as plain text.

<a id="canonical-84c142f013036d7958177b1187fa99510db4b63108d992e1fbb6975440e42a00"></a>

## Next pages — comments / 0c3e6c32a37b / 11

- [comments.attachments_info](data-sources--customer_support_comments--reference--group-001.md#canonical-5b852e4604bfc680b986aa889b3345eca01bd273b515291bf0b8292814052821)
- [Property reference](data-sources--customer_support_comments--reference--group-001.md#canonical-12e9887903e3773a47890b40be28cd522622da1fd04215c1250c156eea0f4f47)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5)

<a id="canonical-5b852e4604bfc680b986aa889b3345eca01bd273b515291bf0b8292814052821"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f9d2d846ed1afa0a48abbb26cd71ed7f5780d5beb139ad301e8caad2f5176be"></a>

## comments.attachments_info — comments.attachments_info / 06ce37af9453 / 2

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5)
- [Property reference](data-sources--customer_support_comments--reference--group-001.md#canonical-12e9887903e3773a47890b40be28cd522622da1fd04215c1250c156eea0f4f47)
- [comments](data-sources--customer_support_comments--reference--group-001.md#canonical-260c3e6a46b7e41abb4555618a3e2f1c1fda9e189cbdfcc95b0a4a9c6f22be19)
- comments.attachments_info

<a id="canonical-07fab5ea44f779dc0baa7cda7a6e064fac789aabdf7d7418b84f848c73d2d711"></a>

Type: `"list"`. Computed.

Information about any attachments (such as screenshots, plain text files) the comment can have.

<a id="canonical-fa354456b500fe375f294cc3915187dd409fa96e1556272b5c3002770d4ba54f"></a>

## Direct properties — comments.attachments_info / 06ce37af9453 / 3

<a id="canonical-da297691a5d0105dbba719a2ee23c0617c560cfa3e8454a07f1530dcdcbfa294"></a>

<a id="canonical-4d0077b06abec7c47084f5cccc57546c596031305c4bebd353037e2a35d81d01"></a>

## attachment property — comments.attachments_info / 06ce37af9453 / 4

Type: `"string"`. Computed.

Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over
HTTP.

<a id="canonical-e9ff3c96a90de2343128e6e098476efab38e13570df423a8431be8423d7c5565"></a>

<a id="canonical-9793ab41850ad6f57a3924de1875a4a0392ea20a8726afe666539d6fca254d6e"></a>

## content_type property — comments.attachments_info / 06ce37af9453 / 5

Type: `"string"`. Computed.

MIME content type of the attachment. Helps the UI to properly display the data.

<a id="canonical-a88615f49e857f5747530cc252357d1f26a90bab38460ea3dde2f276b5b48326"></a>

<a id="canonical-cf06593ee8a738bcccc3448821caa15c53fedeee7638de2a835710b8e3e558eb"></a>

## filename property — comments.attachments_info / 06ce37af9453 / 6

Type: `"string"`. Computed.

Filename of the attachment as provided by the caller.

<a id="canonical-95e09848d5ff9960f093dff4790036c49b40f8988fe1e3a09c21ca3b5de9bc6d"></a>

<a id="canonical-18560ea6e4188f0768f367f001a59a602b81b04afa77baccc235579dc5479ab7"></a>

## tp_id property — comments.attachments_info / 06ce37af9453 / 7

Type: `"string"`. Computed.

Optional ID as assigned by the third-party actually storing the data.

<a id="canonical-594c37664029dc2a88ddbb251706b0e38149be4fa801c713eb29e6c512e86fde"></a>

## Next pages — comments.attachments_info / 06ce37af9453 / 8

- [comments](data-sources--customer_support_comments--reference--group-001.md#canonical-260c3e6a46b7e41abb4555618a3e2f1c1fda9e189cbdfcc95b0a4a9c6f22be19)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5)
