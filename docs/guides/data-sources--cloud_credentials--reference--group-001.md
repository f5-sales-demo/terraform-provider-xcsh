---
page_title: "xcsh_cloud_credentials reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials reference."
---

# xcsh_cloud_credentials reference

<a id="canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-325c8321975e9abd4e647494a3f5561e3b9c3533c70395025a6e96f07022220e"></a>

## Property reference — Property reference / c5753b831306 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- Property reference

<a id="canonical-f1ed88c266b312243b54ad4f211010abcf296842fd4ae83ca02bedb40e8f34a2"></a>

## Direct properties — Property reference / c5753b831306 / 3

<a id="canonical-af333594577c3fe833a2a5ac7e52989f7897fe97f17da686abee5db7d242294f"></a>

<a id="canonical-6f9aec45f06bda98ff00c465ccb11d1acc7983a1efca15d33a6d7ba3c1fb45f3"></a>

## annotations property — Property reference / c5753b831306 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3): complete subsection reference.

- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7): complete subsection reference.

- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3): complete subsection reference.

- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae): complete subsection reference.

<a id="canonical-5136dc2fffd876ef095c58a4993b86d490d04774e2cd0e67a2588c0a42e2e0e3"></a>

<a id="canonical-3cb8660a76144afa5081d2cdaa2d9473e7e7ba249ac797c1823bb6a1f83dbc3e"></a>

## description property — Property reference / c5753b831306 / 5

Type: `"string"`. Computed.

Description of the CloudCredentials.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8): complete subsection reference.

<a id="canonical-ba2d74c16eec9479dafc818a0b360d3d004b3406cf55dc0edd16ae009a6d278e"></a>

<a id="canonical-0bb81035befc071a6abdb4f38660a6445a254c018d4c716d5ecdf1cd2d10953e"></a>

## id property — Property reference / c5753b831306 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0c7346a6962823cb994a0c1579c3b66c92a85e9e11a68a18eba44dfde4d61967"></a>

<a id="canonical-9c8306712bc55dd0b7dad2655ce7251537bfa68fe9136ee3af5db412651d2265"></a>

## labels property — Property reference / c5753b831306 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ee828b8b7fb60827964e2cfc60927539c3833c4b1589c2a72cd11f15e590c3ef"></a>

<a id="canonical-6ed40b101d411a20712b38fc19349788ff7d1f28d291353034c67370bb740a90"></a>

## name property — Property reference / c5753b831306 / 8

Type: `"string"`. Required.

Name of the CloudCredentials.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-680fa2cc9b10e8d2217ed18bb55a18ac1ee06700ef657e7eebc4501b3ad1c146"></a>

<a id="canonical-45335add3284e388c2a881fe4e192ead4b66477906d798cfb82ce0e697f0bb68"></a>

## namespace property — Property reference / c5753b831306 / 9

Type: `"string"`. Required.

Namespace where the CloudCredentials exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-20e86a9130ae3ae4172df66246220e4e1359b181759c8f7637b66a1fa1630aca"></a>

## All schema paths — Property reference / c5753b831306 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_credentials--reference--group-001.md#canonical-af333594577c3fe833a2a5ac7e52989f7897fe97f17da686abee5db7d242294f) |
| `aws_assume_role` | [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-ef2fd2571e3d213b3fb2492afeab56f575e47ec7a9f370ff172df01c1c97d96d) |
| `aws_assume_role.custom_external_id` | [aws_assume_role.custom_external_id](data-sources--cloud_credentials--reference--group-001.md#canonical-78e8372c04c41af2c1854e84cd5fb2c21264568d81cf133619505e5f6a9a797e) |
| `aws_assume_role.duration_seconds` | [aws_assume_role.duration_seconds](data-sources--cloud_credentials--reference--group-001.md#canonical-eec2d29fda60dc3328fcae3d55c0e6f0961d13adf0a937dae4f263f1aa06a20b) |
| `aws_assume_role.external_id_is_optional` | [aws_assume_role.external_id_is_optional](data-sources--cloud_credentials--reference--group-001.md#canonical-8fa10ac6fea62e583bd1e2232e0dbd45c893334f069badf04a1e653649032c6d) |
| `aws_assume_role.external_id_is_tenant_id` | [aws_assume_role.external_id_is_tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-93ecb9cfe1e5e5af17d46925d46893e8291e6a7fdcc59f2c3dbe330259d0105a) |
| `aws_assume_role.role_arn` | [aws_assume_role.role_arn](data-sources--cloud_credentials--reference--group-001.md#canonical-eee977bab3d1241633a13c534f15b6bd807552c05ac86b8a53a41b97e1491c10) |
| `aws_assume_role.session_name` | [aws_assume_role.session_name](data-sources--cloud_credentials--reference--group-001.md#canonical-685ef655c22781d9551bd6a81738f031eb399872536e6a0ab29bfada17db0e53) |
| `aws_assume_role.session_tags` | [aws_assume_role.session_tags](data-sources--cloud_credentials--reference--group-001.md#canonical-d3ac95b9e4a05d319d3032aaaa251f8655c818e92d4a5ab5d80517caae55c245) |
| `aws_secret_key` | [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-eff00760824429095aaeea7f297c9c1a4e30818c03292b8661ebc59f7a29e6e7) |
| `aws_secret_key.access_key` | [aws_secret_key.access_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0b04d95cd47c7533e77f9c17e25eb59c97e89f8d23cb2ffba31be91119028306) |
| `aws_secret_key.secret_key` | [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-51b91ecf23ba138f16eacf8da5e786e317978021d5aee6cd3c34e7018e3604de) |
| `aws_secret_key.secret_key.blindfold_secret_info` | [aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-ed77575afd28ca10d0ab142a38738fc0c990cb6b188913568e6fffd2f7399d8d) |
| `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-2df843b7769db90423e15fac93d21c6e70240f977e8bbdbd480f5dc2092cb19a) |
| `aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_secret_key.secret_key.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-99185e306ad6c3d181fc43e3a622aaa457c820b0a3e87b1b7fb9cae3aaaa0266) |
| `aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_secret_key.secret_key.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-579edc723bdee2e20a533eb30f8121a3a2a83b13f55ae8061baa43ec5ace6197) |
| `aws_secret_key.secret_key.clear_secret_info` | [aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1dd2541066e30172893a7329d5051553ee1c275c4a27ed53049588a5058e3d1b) |
| `aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_secret_key.secret_key.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-11829820d763f9a17d16fd23aa5c9b808b3cadf92dcf64776ea975cbaf7d0522) |
| `aws_secret_key.secret_key.clear_secret_info.url` | [aws_secret_key.secret_key.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-ce3e02336df3519f25b2e6f1e9b7ae053dd84de4cb61a1cf99b7d9f4abca78a5) |
| `azure_client_secret` | [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0314d73bb4aac502558da40d57bb6a6577c675084017026046e6d8da3ca8c786) |
| `azure_client_secret.client_id` | [azure_client_secret.client_id](data-sources--cloud_credentials--reference--group-001.md#canonical-b94cbf60a172cee6312800ea31ea70210eb3ab53764a3a312381b388ed1a2b09) |
| `azure_client_secret.client_secret` | [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-da2d7df6fb21f3bd61d20030d4884a25927064744c4b9f4b67384f2240ee1175) |
| `azure_client_secret.client_secret.blindfold_secret_info` | [azure_client_secret.client_secret.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-a5151879214ebaf0c05ced040b9ca22c6fe84790eeffef2065b61bd7dff87c6b) |
| `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` | [azure_client_secret.client_secret.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-ea5b3de9cedc11d374daaae7eadf2e47a45cfad9a99458273f442c36d1690617) |
| `azure_client_secret.client_secret.blindfold_secret_info.location` | [azure_client_secret.client_secret.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-b97dec39d07ed8cca0cd06ba0fc65c68af6c396a0d8bb591edb28e5bf7d15c73) |
| `azure_client_secret.client_secret.blindfold_secret_info.store_provider` | [azure_client_secret.client_secret.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-8ae894dddce5a3eb1db0e2ec5bb45b3e559aa98f633416d5af23dad58003a4b4) |
| `azure_client_secret.client_secret.clear_secret_info` | [azure_client_secret.client_secret.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-f9835e15fe80b470daa8e09e1e9e5da70cdcab305163830a798b644250cca049) |
| `azure_client_secret.client_secret.clear_secret_info.provider_ref` | [azure_client_secret.client_secret.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-f6e3a16999ede3b035ada3e18ba77032f3f7fadcc8b7539b57ec0b647a038b0b) |
| `azure_client_secret.client_secret.clear_secret_info.url` | [azure_client_secret.client_secret.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-5ee33df2708055f1543b22aa02df846e30dfca259773c056c10a02496546aea3) |
| `azure_client_secret.subscription_id` | [azure_client_secret.subscription_id](data-sources--cloud_credentials--reference--group-001.md#canonical-9b59b007f01342256d654e8f4bc512f2573fd804c1d200411812bad57e3627ee) |
| `azure_client_secret.tenant_id` | [azure_client_secret.tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-abf21a880111721d6f8a00063534dad10aed66581eb4c0f4da75c3d410f4102d) |
| `azure_pfx_certificate` | [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-1869ed338f95c519811afaf9a7d11e55781aa60e95ac17a165cdb9dd90f852aa) |
| `azure_pfx_certificate.certificate_url` | [azure_pfx_certificate.certificate_url](data-sources--cloud_credentials--reference--group-001.md#canonical-8042a22b43e996035fd45bb210dad650890c708cfbb4b14d98a765aabe7c99cf) |
| `azure_pfx_certificate.client_id` | [azure_pfx_certificate.client_id](data-sources--cloud_credentials--reference--group-001.md#canonical-2c75704ac70ccb40effe7361f10bffe73851eaf09b4ba4429206e055f4130133) |
| `azure_pfx_certificate.password` | [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-1402737703f2daaa36872ed0034fc8e7da49dfe70912405533a9df705be571e1) |
| `azure_pfx_certificate.password.blindfold_secret_info` | [azure_pfx_certificate.password.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-2822eeb17419316c40b1f35d331e32556aa5a2fa153a8c3cd389038e61595033) |
| `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` | [azure_pfx_certificate.password.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-d7573e2b759dba7efc076198592ac02564c4742a37b4041c3b9f8fd12db544c4) |
| `azure_pfx_certificate.password.blindfold_secret_info.location` | [azure_pfx_certificate.password.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-d81e0334bfe0d0d3a74531e5b667547153c91bb61f37ed2137a918b02aef4739) |
| `azure_pfx_certificate.password.blindfold_secret_info.store_provider` | [azure_pfx_certificate.password.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-de5ca7daba62fec4ad78a1c8db2c2982855ec25aa3ad466427e6562b6d8c74f4) |
| `azure_pfx_certificate.password.clear_secret_info` | [azure_pfx_certificate.password.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-59f28889a8ff1ffa51ea627f4271b8db2d8076de28a38859c74949886c6106ae) |
| `azure_pfx_certificate.password.clear_secret_info.provider_ref` | [azure_pfx_certificate.password.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-f82f4e0b9ec95db221e72c757544adba78f31151c0f03bd8eae59aad1bb7c7a0) |
| `azure_pfx_certificate.password.clear_secret_info.url` | [azure_pfx_certificate.password.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-dc8315289ca40913b0d9a73d0d629bc031a7afbacfc825b238af3104bcfc61d0) |
| `azure_pfx_certificate.subscription_id` | [azure_pfx_certificate.subscription_id](data-sources--cloud_credentials--reference--group-001.md#canonical-de74aaa755ad9f076fc1dda86380782bc03b375cfe0beac4d8a866e1dd0035d5) |
| `azure_pfx_certificate.tenant_id` | [azure_pfx_certificate.tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-70993c4eae956d8a2525d227f456b89401c27b78139b2ebfe3ef66202ff9821e) |
| `description` | [description](data-sources--cloud_credentials--reference--group-001.md#canonical-5136dc2fffd876ef095c58a4993b86d490d04774e2cd0e67a2588c0a42e2e0e3) |
| `gcp_cred_file` | [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-07cca961171f33df139a5f250fabdca0404e79eb8d46f6ebb178549db900ebb2) |
| `gcp_cred_file.credential_file` | [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-3054c970ce52a6d8cf2a0d9f9ad4cdff9fbe549f4a4a3c9521beb03e671442b1) |
| `gcp_cred_file.credential_file.blindfold_secret_info` | [gcp_cred_file.credential_file.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-0f2e7d7dd4c5b46d8587254c0d2322f45b6da5e9fd5c4572a9a988455ecfd5e6) |
| `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-6289febd13613ab53d388708b1bc1372a00d19f4455f90e8df3a69b6f53ec92e) |
| `gcp_cred_file.credential_file.blindfold_secret_info.location` | [gcp_cred_file.credential_file.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-f64bb346d40676d9db79da34f6edad0ba90c4d05950a4f465cfd9f1fcdd21e9a) |
| `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-b8ed2d6cfdc48c146b10dd2cd1e09ed98ffb4cba7d7516ee99c3ec6ece69ac71) |
| `gcp_cred_file.credential_file.clear_secret_info` | [gcp_cred_file.credential_file.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-81f180237e2234202289bba78b63887dfbb71fb52d5b7de5b5c336b81c78bff4) |
| `gcp_cred_file.credential_file.clear_secret_info.provider_ref` | [gcp_cred_file.credential_file.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-07c7d7c7052c3e7058d871ed61c45abd550b235e5f3cc1d1b637a119ce813e3f) |
| `gcp_cred_file.credential_file.clear_secret_info.url` | [gcp_cred_file.credential_file.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-4236672525f9008e075c97b073b862b9a9609ef5e5ab7f3fadc3c4f7ba136114) |
| `id` | [id](data-sources--cloud_credentials--reference--group-001.md#canonical-ba2d74c16eec9479dafc818a0b360d3d004b3406cf55dc0edd16ae009a6d278e) |
| `labels` | [labels](data-sources--cloud_credentials--reference--group-001.md#canonical-0c7346a6962823cb994a0c1579c3b66c92a85e9e11a68a18eba44dfde4d61967) |
| `name` | [name](data-sources--cloud_credentials--reference--group-001.md#canonical-ee828b8b7fb60827964e2cfc60927539c3833c4b1589c2a72cd11f15e590c3ef) |
| `namespace` | [namespace](data-sources--cloud_credentials--reference--group-001.md#canonical-680fa2cc9b10e8d2217ed18bb55a18ac1ee06700ef657e7eebc4501b3ad1c146) |

<a id="canonical-3a6b607f89ccae2dbb45d4932f76f74df499d46b320cf48d7af0db1a8304206d"></a>

## Next pages — Property reference / c5753b831306 / 11

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff70ae8be3b43db55a93e42cadb57aa044975f5fd67d352619dd0f8d92516c5c"></a>

## aws_assume_role — aws_assume_role / 23cca173feee / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- aws_assume_role

<a id="canonical-ef2fd2571e3d213b3fb2492afeab56f575e47ec7a9f370ff172df01c1c97d96d"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_assume\_role, aws\_secret\_key, azure\_client\_secret, azure\_pfx\_certificate,
gcp\_cred\_file\] AWS Assume Role to Handle Delegated Access.

Upstream description:

AWS Assume Role to Handle Delegated Access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-external_id": "[\"custom_external_id\",\"external_id_is_optional\",\"external_id_is_tenant_id\"]"
}
```

OneOf alternatives in this subsection:

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-ef2fd2571e3d213b3fb2492afeab56f575e47ec7a9f370ff172df01c1c97d96d)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-eff00760824429095aaeea7f297c9c1a4e30818c03292b8661ebc59f7a29e6e7)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0314d73bb4aac502558da40d57bb6a6577c675084017026046e6d8da3ca8c786)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-1869ed338f95c519811afaf9a7d11e55781aa60e95ac17a165cdb9dd90f852aa)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-07cca961171f33df139a5f250fabdca0404e79eb8d46f6ebb178549db900ebb2)

Select alternatives according to the provider validators above.

<a id="canonical-e0e46e808ca56484f1e84556d27e73d577035f4e6a702594eb297cb371973503"></a>

## Direct properties — aws_assume_role / 23cca173feee / 3

<a id="canonical-78e8372c04c41af2c1854e84cd5fb2c21264568d81cf133619505e5f6a9a797e"></a>

<a id="canonical-0b65e5f8f92fc424023cad1cf31bfdb72497c89a86a2ee8c14de172bd6bf70ed"></a>

## custom_external_id property — aws_assume_role / 23cca173feee / 4

Type: `"string"`. Computed.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  }
}
```

<a id="canonical-eec2d29fda60dc3328fcae3d55c0e6f0961d13adf0a937dae4f263f1aa06a20b"></a>

<a id="canonical-8c07102d0e3be2554f5b2a5be76881711a63d577deca9a17c90ca6e5a96ed1d3"></a>

## duration_seconds property — aws_assume_role / 23cca173feee / 5

Type: `"number"`. Computed.

The duration, in seconds of the role session.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

- [external_id_is_optional](data-sources--cloud_credentials--reference--group-001.md#canonical-55d984242b3064838295650a6b3144ab5781b0e17b2f1240623055b89dd96d93): complete subsection reference.

- [external_id_is_tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-d62112aed1b14997b956bb6312291aec14b1ff91059ba0deb5ac937e1022b7b1): complete subsection reference.

<a id="canonical-eee977bab3d1241633a13c534f15b6bd807552c05ac86b8a53a41b97e1491c10"></a>

<a id="canonical-950434be3f828ef8d0c18db0a810caae05bc91e1cc073d4f0f63f91601230e78"></a>

## role_arn property — aws_assume_role / 23cca173feee / 6

Type: `"string"`. Computed.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20,
    "pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  }
}
```

<a id="canonical-685ef655c22781d9551bd6a81738f031eb399872536e6a0ab29bfada17db0e53"></a>

<a id="canonical-5fdd36856af96631d8bc78ce9f578b7f65576f7cba94d3167f13e9a8c0b83aae"></a>

## session_name property — aws_assume_role / 23cca173feee / 7

Type: `"string"`. Computed.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2,
    "pattern": "[\\\\w+=,.@-]*"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  }
}
```

<a id="canonical-d3ac95b9e4a05d319d3032aaaa251f8655c818e92d4a5ab5d80517caae55c245"></a>

<a id="canonical-7bffd9d9dbedabb0e80c9e27c3d43f9d00cde8a43c009c97a9f0d3b1c87d9481"></a>

## session_tags property — aws_assume_role / 23cca173feee / 8

Type: `["map", "string"]`. Computed.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-e4449d1f521bb2a51e08433193038053215becc077b091391874a4cfc144a967"></a>

## Next pages — aws_assume_role / 23cca173feee / 9

- [aws_assume_role.external_id_is_optional](data-sources--cloud_credentials--reference--group-001.md#canonical-55d984242b3064838295650a6b3144ab5781b0e17b2f1240623055b89dd96d93)
- [aws_assume_role.external_id_is_tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-d62112aed1b14997b956bb6312291aec14b1ff91059ba0deb5ac937e1022b7b1)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-55d984242b3064838295650a6b3144ab5781b0e17b2f1240623055b89dd96d93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15cf983985d98cfc880b1fea1fd4783fe132006ab071b2ab27948cdc901cb901"></a>

## aws_assume_role.external_id_is_optional — aws_assume_role.external_id_is_optional / 39ef936f8330 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3)
- aws_assume_role.external_id_is_optional

<a id="canonical-8fa10ac6fea62e583bd1e2232e0dbd45c893334f069badf04a1e653649032c6d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external id is optional.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5a9eb13f9b7acd38a3dfbfc4b21456324a92e0cabcc4366de0ef11c9ce573375"></a>

## Direct properties — aws_assume_role.external_id_is_optional / 39ef936f8330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-262812e6f584195bd4e1632fe64797228ab36448f33d90ecf539be4c6ca1b198"></a>

## Next pages — aws_assume_role.external_id_is_optional / 39ef936f8330 / 4

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-d62112aed1b14997b956bb6312291aec14b1ff91059ba0deb5ac937e1022b7b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f1f5a7c02342c1612d2eb0bb57c32aa80ab2b12f7100a92b3436bb4055b1eb"></a>

## aws_assume_role.external_id_is_tenant_id — aws_assume_role.external_id_is_tenant_id / b9bf090a4888 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3)
- aws_assume_role.external_id_is_tenant_id

<a id="canonical-93ecb9cfe1e5e5af17d46925d46893e8291e6a7fdcc59f2c3dbe330259d0105a"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9ea300cbc79b1fdb7a0981f25f9a7a74818d1e32bdad8640c77afd5d9c3423b1"></a>

## Direct properties — aws_assume_role.external_id_is_tenant_id / b9bf090a4888 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6108e541869ca4871ae52134ad4ded0253d3c2a931571d07a89aed2616259af3"></a>

## Next pages — aws_assume_role.external_id_is_tenant_id / b9bf090a4888 / 4

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-e729d9d7b19152edb66aa92d45c5f6b1d09513a9e05e250edfe8732791d590b3)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd4542699a6857b5b4c8a25dd10d64d8ab51b03f96744f24faf5da999a7383cf"></a>

## aws_secret_key — aws_secret_key / d60d5e3b8eee / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- aws_secret_key

<a id="canonical-eff00760824429095aaeea7f297c9c1a4e30818c03292b8661ebc59f7a29e6e7"></a>

Type: `"single"`. Computed.

AWS Programmatic Access Credentials type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d14ba2539aa126503e15b83da7e4d1b2f3ee4ea0567adbaf2e796766380c461a"></a>

## Direct properties — aws_secret_key / d60d5e3b8eee / 3

<a id="canonical-0b04d95cd47c7533e77f9c17e25eb59c97e89f8d23cb2ffba31be91119028306"></a>

<a id="canonical-655ba4c94d1971568895b8945024313d1b5a512830680fcc73a3feac39f7ef9b"></a>

## access_key property — aws_secret_key / d60d5e3b8eee / 4

Type: `"string"`. Computed.

Access Key ID. Access key ID for your AWS account.

Upstream description:

Access key ID for your AWS account.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759): complete subsection reference.

<a id="canonical-4e81c18d34de455364ccd615c964f5b8e485e268bcacb99cc24b4ef88e73cda7"></a>

## Next pages — aws_secret_key / d60d5e3b8eee / 5

- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9814bae2c421a84d2ce1faa14c96174af0c84c621aff8f3a3e325f9eb16d08df"></a>

## aws_secret_key.secret_key — aws_secret_key.secret_key / 6afa52f54fc3 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7)
- aws_secret_key.secret_key

<a id="canonical-51b91ecf23ba138f16eacf8da5e786e317978021d5aee6cd3c34e7018e3604de"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-d6baaf90b139a2a988cfca3e43066070254a380ae5b192a6017b9f06d230c7cd"></a>

## Direct properties — aws_secret_key.secret_key / 6afa52f54fc3 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-b3b3a9c7dd88e044813abd6e3164aa08706814ccb7be9b327adf2249a5eda8de): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-43c7aa272308be8a8d6d5a02eaa5042185cbf9e876e03ac35a29b84b42265fd2): complete subsection reference.

<a id="canonical-f2c0c340cf5257e785b386daed0c270e527621a4781186622cd4a9e610186c7a"></a>

## Next pages — aws_secret_key.secret_key / 6afa52f54fc3 / 4

- [aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-b3b3a9c7dd88e044813abd6e3164aa08706814ccb7be9b327adf2249a5eda8de)
- [aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-43c7aa272308be8a8d6d5a02eaa5042185cbf9e876e03ac35a29b84b42265fd2)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-b3b3a9c7dd88e044813abd6e3164aa08706814ccb7be9b327adf2249a5eda8de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-617e7c9fad46ccea8fb112247a6fc88601d68db5e80d684dcfe4b45ac5bed416"></a>

## aws_secret_key.secret_key.blindfold_secret_info — aws_secret_key.secret_key.blindfold_secret_info / 5ffa6cbd6a08 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7)
- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759)
- aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-ed77575afd28ca10d0ab142a38738fc0c990cb6b188913568e6fffd2f7399d8d"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-81abb1328d209178f22ba0428d9d55b0ed02e052174297f9273f3b8a221603e5"></a>

## Direct properties — aws_secret_key.secret_key.blindfold_secret_info / 5ffa6cbd6a08 / 3

<a id="canonical-2df843b7769db90423e15fac93d21c6e70240f977e8bbdbd480f5dc2092cb19a"></a>

<a id="canonical-239dea039aad6ec8b9bb70e5d2d5050d7bef846dc230c9ea14f78a608df78340"></a>

## decryption_provider property — aws_secret_key.secret_key.blindfold_secret_info / 5ffa6cbd6a08 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-99185e306ad6c3d181fc43e3a622aaa457c820b0a3e87b1b7fb9cae3aaaa0266"></a>

<a id="canonical-0bd5108bd2a12da15d73287bf9a19a6f50f297e0df983b189899b6dd728ce8a3"></a>

## location property — aws_secret_key.secret_key.blindfold_secret_info / 5ffa6cbd6a08 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-579edc723bdee2e20a533eb30f8121a3a2a83b13f55ae8061baa43ec5ace6197"></a>

<a id="canonical-9550eaa422eaf0db48e1263359fe9bfe258c80c6df420c5c7b09c07eaaa61932"></a>

## store_provider property — aws_secret_key.secret_key.blindfold_secret_info / 5ffa6cbd6a08 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-34b9a162c30655aacec02992aa6c04f96567b15d30f2ccf4161a6da990a18c7f"></a>

## Next pages — aws_secret_key.secret_key.blindfold_secret_info / 5ffa6cbd6a08 / 7

- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-43c7aa272308be8a8d6d5a02eaa5042185cbf9e876e03ac35a29b84b42265fd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d91c2188c3a9d06d0b896f1dec87fba497f858cd79728bb619233166ba1596a2"></a>

## aws_secret_key.secret_key.clear_secret_info — aws_secret_key.secret_key.clear_secret_info / ba9749c9fee7 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-9218282a7e2c81d7f41843b736f78ae638085b02d1cda698590eb14a0ce55ee7)
- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759)
- aws_secret_key.secret_key.clear_secret_info

<a id="canonical-1dd2541066e30172893a7329d5051553ee1c275c4a27ed53049588a5058e3d1b"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-55f24dd30c21f709f2608aa31e04fff7ccec2d9d19ece79b91581bbe35b2a651"></a>

## Direct properties — aws_secret_key.secret_key.clear_secret_info / ba9749c9fee7 / 3

<a id="canonical-11829820d763f9a17d16fd23aa5c9b808b3cadf92dcf64776ea975cbaf7d0522"></a>

<a id="canonical-271f5c3fdf87026d4eee41ed2b29cad5fb3d7399f8517cc5a4748c00b74d4deb"></a>

## provider_ref property — aws_secret_key.secret_key.clear_secret_info / ba9749c9fee7 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ce3e02336df3519f25b2e6f1e9b7ae053dd84de4cb61a1cf99b7d9f4abca78a5"></a>

<a id="canonical-c90dd5720ddc51d646e3e5dab43dd39acf3ca8341e1816cce2be6518b93adbc1"></a>

## url property — aws_secret_key.secret_key.clear_secret_info / ba9749c9fee7 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-bf8474eb6d6477ead149eb6758c77680a6d34ce8a8acea2d0950ab0b8d332d43"></a>

## Next pages — aws_secret_key.secret_key.clear_secret_info / ba9749c9fee7 / 6

- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-00331de55298db302ebf67ba1e63fcbbbb18bb5e98f1d7353baa8ece54742759)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e24844f39df04b59641fd56dd487eb81b3e0b39ae690d3a7aefd674e1d468d9"></a>

## azure_client_secret — azure_client_secret / 14ce1334e2d1 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- azure_client_secret

<a id="canonical-0314d73bb4aac502558da40d57bb6a6577c675084017026046e6d8da3ca8c786"></a>

Type: `"single"`. Computed.

Azure Client Secret. Azure Credentials Client Secret type.

Upstream description:

Azure Credentials Client Secret type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-772e476dd5e60f6fbc63bce99d22c3f6dc3aebe3fe539f57511519ea33662680"></a>

## Direct properties — azure_client_secret / 14ce1334e2d1 / 3

<a id="canonical-b94cbf60a172cee6312800ea31ea70210eb3ab53764a3a312381b388ed1a2b09"></a>

<a id="canonical-bfb5d6cba43f7bc94d34b8a76e0cc6b821d4d63e125b99dacfd0a2d78f1ae414"></a>

## client_id property — azure_client_secret / 14ce1334e2d1 / 4

Type: `"string"`. Computed.

Client ID for your Azure service principal.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750): complete subsection reference.

<a id="canonical-9b59b007f01342256d654e8f4bc512f2573fd804c1d200411812bad57e3627ee"></a>

<a id="canonical-f516c060c4fa697cfbc3713622314ed3328bb818816bc399bfb3ad2412b33051"></a>

## subscription_id property — azure_client_secret / 14ce1334e2d1 / 5

Type: `"string"`. Computed.

Subscription ID for your Azure service principal.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-abf21a880111721d6f8a00063534dad10aed66581eb4c0f4da75c3d410f4102d"></a>

<a id="canonical-679623d11fd0a4c723d570ac8a8fd18f1855ffc9a50fb0264b8795dc555cfd49"></a>

## tenant_id property — azure_client_secret / 14ce1334e2d1 / 6

Type: `"string"`. Computed.

Tenant ID for your Azure service principal.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-bd638daa563571391aecd6c92e0fb6ef5782bf639d287609bf26722a02a86fbb"></a>

## Next pages — azure_client_secret / 14ce1334e2d1 / 7

- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53a47f7ae4948e31a3bd6eba7d07b321db78828c0ea675a0f4cb533aa3361ea8"></a>

## azure_client_secret.client_secret — azure_client_secret.client_secret / 1e40af47e765 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3)
- azure_client_secret.client_secret

<a id="canonical-da2d7df6fb21f3bd61d20030d4884a25927064744c4b9f4b67384f2240ee1175"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-7633515462a6a751a3af1d863e11328135546ffbc59d3f35e93decf8639a892a"></a>

## Direct properties — azure_client_secret.client_secret / 1e40af47e765 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-74059f56d3f2604e367c9134cf1d433877e8407bb47c389dff712c21daa6ff47): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-7da6ff76568582cec0a74a2caf1d4f250819c82fb739064cf40ab404d4f4dfa1): complete subsection reference.

<a id="canonical-c7d4f8e0c053c7ba321c485fd045191aaae01666ae8e7550d4b0658415bda8a2"></a>

## Next pages — azure_client_secret.client_secret / 1e40af47e765 / 4

- [azure_client_secret.client_secret.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-74059f56d3f2604e367c9134cf1d433877e8407bb47c389dff712c21daa6ff47)
- [azure_client_secret.client_secret.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-7da6ff76568582cec0a74a2caf1d4f250819c82fb739064cf40ab404d4f4dfa1)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-74059f56d3f2604e367c9134cf1d433877e8407bb47c389dff712c21daa6ff47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a59c747aac6d04b84e62bd3a7419d35683e7dc8e9d828b0655511e566358a07"></a>

## azure_client_secret.client_secret.blindfold_secret_info — azure_client_secret.client_secret.blindfold_secret_info / 38138bceee49 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3)
- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750)
- azure_client_secret.client_secret.blindfold_secret_info

<a id="canonical-a5151879214ebaf0c05ced040b9ca22c6fe84790eeffef2065b61bd7dff87c6b"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-71f764c9f2ddca142dc60db75cad6a37999fd8c1e81d4a848c3140ed421cdfe4"></a>

## Direct properties — azure_client_secret.client_secret.blindfold_secret_info / 38138bceee49 / 3

<a id="canonical-ea5b3de9cedc11d374daaae7eadf2e47a45cfad9a99458273f442c36d1690617"></a>

<a id="canonical-0bcac739237c12ca4bb7527acb8455d0849a2b4365a334807d2ea32af1fd80dc"></a>

## decryption_provider property — azure_client_secret.client_secret.blindfold_secret_info / 38138bceee49 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b97dec39d07ed8cca0cd06ba0fc65c68af6c396a0d8bb591edb28e5bf7d15c73"></a>

<a id="canonical-47d702826229ef848417c41832cde4771ad52d0eb2bf4798a104891a26a52905"></a>

## location property — azure_client_secret.client_secret.blindfold_secret_info / 38138bceee49 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-8ae894dddce5a3eb1db0e2ec5bb45b3e559aa98f633416d5af23dad58003a4b4"></a>

<a id="canonical-632dfee0e83605b22cdbdbd09841e359c94bcf779a933e64897621b8052a31f4"></a>

## store_provider property — azure_client_secret.client_secret.blindfold_secret_info / 38138bceee49 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2c7640f54fa0f1b8c608d8a9b85b485f00d0ac75d2a40d1a5bf4db125efa74e1"></a>

## Next pages — azure_client_secret.client_secret.blindfold_secret_info / 38138bceee49 / 7

- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-7da6ff76568582cec0a74a2caf1d4f250819c82fb739064cf40ab404d4f4dfa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-464eb7561eecb2f6547d8ad51b17de95d8b7409b56e6af7264d25527148c6dfa"></a>

## azure_client_secret.client_secret.clear_secret_info — azure_client_secret.client_secret.clear_secret_info / afa10edb9b3b / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-07f5333599b090c4ed62878bdad7479fcf068220921b63a4765beed0770712e3)
- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750)
- azure_client_secret.client_secret.clear_secret_info

<a id="canonical-f9835e15fe80b470daa8e09e1e9e5da70cdcab305163830a798b644250cca049"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-71b2e206a8e6d9c1b6419f14c0ae11d01750b4501cefe8e91c4d3f0db4476785"></a>

## Direct properties — azure_client_secret.client_secret.clear_secret_info / afa10edb9b3b / 3

<a id="canonical-f6e3a16999ede3b035ada3e18ba77032f3f7fadcc8b7539b57ec0b647a038b0b"></a>

<a id="canonical-35744457d9ef1a1a88611a02cc08e16c9982442da96a9890f9105388d7628fa7"></a>

## provider_ref property — azure_client_secret.client_secret.clear_secret_info / afa10edb9b3b / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-5ee33df2708055f1543b22aa02df846e30dfca259773c056c10a02496546aea3"></a>

<a id="canonical-53f1817afc8e1c4600f8dc157cc84bb8a64ddd9d52e720dafbb96f29419d9555"></a>

## url property — azure_client_secret.client_secret.clear_secret_info / afa10edb9b3b / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-783201b9ef391f673cde7d2d5fd1ee2853e9f32eaab769c11caf8fe5331a1562"></a>

## Next pages — azure_client_secret.client_secret.clear_secret_info / afa10edb9b3b / 6

- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-7e6fe01e4974b877eb9cf1f44dcd138a4f533b7298297e755046d238836c5750)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55ab792d854547cd7315563d76d2cb34da7463dbf72ff2c8065c4eb53437888c"></a>

## azure_pfx_certificate — azure_pfx_certificate / a09b20a61c64 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- azure_pfx_certificate

<a id="canonical-1869ed338f95c519811afaf9a7d11e55781aa60e95ac17a165cdb9dd90f852aa"></a>

Type: `"single"`. Computed.

Azure Credentials Client Certificate type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c70b7fd690b89140051cc6835b94193173eda4105d971e5223ecf3566bc71af4"></a>

## Direct properties — azure_pfx_certificate / a09b20a61c64 / 3

<a id="canonical-8042a22b43e996035fd45bb210dad650890c708cfbb4b14d98a765aabe7c99cf"></a>

<a id="canonical-ea86bd65f022ed0685308828d4e781277f290f0a1446a5bcb62d3f7f971b9845"></a>

## certificate_url property — azure_pfx_certificate / a09b20a61c64 / 4

Type: `"string"`. Computed.

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Upstream description:

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2c75704ac70ccb40effe7361f10bffe73851eaf09b4ba4429206e055f4130133"></a>

<a id="canonical-e4d2a09dcfe9c0697b8744336ace10770412fbad890ae12d1747557d8ef0c456"></a>

## client_id property — azure_pfx_certificate / a09b20a61c64 / 5

Type: `"string"`. Computed.

Client ID for your Azure service principal.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [password](data-sources--cloud_credentials--reference--group-001.md#canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574): complete subsection reference.

<a id="canonical-de74aaa755ad9f076fc1dda86380782bc03b375cfe0beac4d8a866e1dd0035d5"></a>

<a id="canonical-b11953ff1fa36b2641159649aed1e4a22737996ffc5729e8aca4537811aad3b3"></a>

## subscription_id property — azure_pfx_certificate / a09b20a61c64 / 6

Type: `"string"`. Computed.

Subscription ID for your Azure service principal.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-70993c4eae956d8a2525d227f456b89401c27b78139b2ebfe3ef66202ff9821e"></a>

<a id="canonical-2d8a8fbf6aaafd4dab24075594c86c111fdbcdf3574c50c5ead8a8bcb24ff7fe"></a>

## tenant_id property — azure_pfx_certificate / a09b20a61c64 / 7

Type: `"string"`. Computed.

Tenant ID for your Azure service principal.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-98e2e73594820824f6d608fc4843fff5e89fbea861782af34397f6625eae8e87"></a>

## Next pages — azure_pfx_certificate / a09b20a61c64 / 8

- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc060200281b52790ff4f32aa2fa947f5d6ce07c18456819774726376933f2ac"></a>

## azure_pfx_certificate.password — azure_pfx_certificate.password / 7aaa849fc9d2 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae)
- azure_pfx_certificate.password

<a id="canonical-1402737703f2daaa36872ed0034fc8e7da49dfe70912405533a9df705be571e1"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-ec9a5ebdb063703ccc2cec6b1372601a6ec086cce628d4a5733a3a2295a3dc4e"></a>

## Direct properties — azure_pfx_certificate.password / 7aaa849fc9d2 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-6680c52ba16dad1ed2637f84b1a8f1c7f18366ebeaf88375cb569e897f4df388): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-61681fe7b0efce631bca5f60a8a7a557ee0abfdc44bffaac4854ad0ec937cf9e): complete subsection reference.

<a id="canonical-02f1f089d73d1facead73ef728257bdce488510274f79b71d36913f4e562b27b"></a>

## Next pages — azure_pfx_certificate.password / 7aaa849fc9d2 / 4

- [azure_pfx_certificate.password.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-6680c52ba16dad1ed2637f84b1a8f1c7f18366ebeaf88375cb569e897f4df388)
- [azure_pfx_certificate.password.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-61681fe7b0efce631bca5f60a8a7a557ee0abfdc44bffaac4854ad0ec937cf9e)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-6680c52ba16dad1ed2637f84b1a8f1c7f18366ebeaf88375cb569e897f4df388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8979a94f99943eeb241c53da96242aae0c5f5e4895e7514714465bea77ce47a5"></a>

## azure_pfx_certificate.password.blindfold_secret_info — azure_pfx_certificate.password.blindfold_secret_info / abba5c869a8a / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae)
- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574)
- azure_pfx_certificate.password.blindfold_secret_info

<a id="canonical-2822eeb17419316c40b1f35d331e32556aa5a2fa153a8c3cd389038e61595033"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-30f43ade37e74bfd39b071d5322224e824fde241d6956cd628408140ef0305eb"></a>

## Direct properties — azure_pfx_certificate.password.blindfold_secret_info / abba5c869a8a / 3

<a id="canonical-d7573e2b759dba7efc076198592ac02564c4742a37b4041c3b9f8fd12db544c4"></a>

<a id="canonical-bc12f9491b3140281794f7538ff4adb95ea76e2d902e16ed87f6dcd279fb661f"></a>

## decryption_provider property — azure_pfx_certificate.password.blindfold_secret_info / abba5c869a8a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d81e0334bfe0d0d3a74531e5b667547153c91bb61f37ed2137a918b02aef4739"></a>

<a id="canonical-4f6df6103f43c52416061753200cc8243b089285e5abcd1fd103763e079d226f"></a>

## location property — azure_pfx_certificate.password.blindfold_secret_info / abba5c869a8a / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-de5ca7daba62fec4ad78a1c8db2c2982855ec25aa3ad466427e6562b6d8c74f4"></a>

<a id="canonical-8dbae8789d5462d1042eaf307fcea7cd2c484130a65576afa05b7a7307412587"></a>

## store_provider property — azure_pfx_certificate.password.blindfold_secret_info / abba5c869a8a / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a0dbf89a9b087df5a14f6a22e1f11542f0ffc28d69a7576632a297c071bd44b3"></a>

## Next pages — azure_pfx_certificate.password.blindfold_secret_info / abba5c869a8a / 7

- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-61681fe7b0efce631bca5f60a8a7a557ee0abfdc44bffaac4854ad0ec937cf9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5edcfad3049d7690441cfb02c540bc34faccfa65db93db0c05100c1124491f10"></a>

## azure_pfx_certificate.password.clear_secret_info — azure_pfx_certificate.password.clear_secret_info / c0e5bc6f71d8 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-9734b8dad3dabf19b525aa639c70123358001c369fd8b2eb8d7ba08f9d1fc4ae)
- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574)
- azure_pfx_certificate.password.clear_secret_info

<a id="canonical-59f28889a8ff1ffa51ea627f4271b8db2d8076de28a38859c74949886c6106ae"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ea665834602d7a638f90d0cb58615722b7f0717dc9429b662b5287e3426b67f2"></a>

## Direct properties — azure_pfx_certificate.password.clear_secret_info / c0e5bc6f71d8 / 3

<a id="canonical-f82f4e0b9ec95db221e72c757544adba78f31151c0f03bd8eae59aad1bb7c7a0"></a>

<a id="canonical-d6601cfbe5e414781dcef080e33904f6ff261fc3048d047c55cc7c1baeb74739"></a>

## provider_ref property — azure_pfx_certificate.password.clear_secret_info / c0e5bc6f71d8 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-dc8315289ca40913b0d9a73d0d629bc031a7afbacfc825b238af3104bcfc61d0"></a>

<a id="canonical-adc70f601f85abad880d3b67ebb99b330b387c7964ee77a362dce25c5bdc7ce1"></a>

## url property — azure_pfx_certificate.password.clear_secret_info / c0e5bc6f71d8 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-7466cfba2e77039cb448f06a1a69aff8892ca60362a89f2762008893e0e07502"></a>

## Next pages — azure_pfx_certificate.password.clear_secret_info / c0e5bc6f71d8 / 6

- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-639a6377750191b6eff95fdc1f51ec471c24f7b41e1d0641a143cc711e2ad574)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4404b5debe152f8d238bfdecb75366e178522bdd6977f0bd77512b7dd299b377"></a>

## gcp_cred_file — gcp_cred_file / 62084f1a49eb / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- gcp_cred_file

<a id="canonical-07cca961171f33df139a5f250fabdca0404e79eb8d46f6ebb178549db900ebb2"></a>

Type: `"single"`. Computed.

Configuration parameter for gcp cred file.

Upstream description:

GCP Credentials type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-15e36b47f16900c25a67cd25da995e0598c4c4ad6587ce66515b9ce99653d4ee"></a>

## Direct properties — gcp_cred_file / 62084f1a49eb / 3

- [credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221): complete subsection reference.

<a id="canonical-dfc385725bd812243d23b48b0436fdb4928ed5da3367a7c192b56c290daa3b8e"></a>

## Next pages — gcp_cred_file / 62084f1a49eb / 4

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb219b44f62f502491a63f22817c719140790ee9283b0f7b6a93e136c33cc81c"></a>

## gcp_cred_file.credential_file — gcp_cred_file.credential_file / 4cf9ebf3ec5c / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8)
- gcp_cred_file.credential_file

<a id="canonical-3054c970ce52a6d8cf2a0d9f9ad4cdff9fbe549f4a4a3c9521beb03e671442b1"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-5c2b504e481c8d801cb877a0738783e898d33427029592d395148a605c48f197"></a>

## Direct properties — gcp_cred_file.credential_file / 4cf9ebf3ec5c / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-cca759b7fc1bf597c7eb97e8546a0530f4748af5524963f862b01bc8dd73c188): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-ff665e06caa8b4391b73df4d1c574f32cb86ff5d2fc474f35611593749454c1a): complete subsection reference.

<a id="canonical-3db2d086a0fbf186f65891cbcc416342f2818bd2135172947d37fb55ebbee204"></a>

## Next pages — gcp_cred_file.credential_file / 4cf9ebf3ec5c / 4

- [gcp_cred_file.credential_file.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-cca759b7fc1bf597c7eb97e8546a0530f4748af5524963f862b01bc8dd73c188)
- [gcp_cred_file.credential_file.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-ff665e06caa8b4391b73df4d1c574f32cb86ff5d2fc474f35611593749454c1a)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-cca759b7fc1bf597c7eb97e8546a0530f4748af5524963f862b01bc8dd73c188"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7cf80d0859566bb2f70205960195a2eba7712316600bd7527d5ee6aa8f5079e"></a>

## gcp_cred_file.credential_file.blindfold_secret_info — gcp_cred_file.credential_file.blindfold_secret_info / cf8bec656890 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8)
- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221)
- gcp_cred_file.credential_file.blindfold_secret_info

<a id="canonical-0f2e7d7dd4c5b46d8587254c0d2322f45b6da5e9fd5c4572a9a988455ecfd5e6"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-60bb1687986dd9acd07300dfa9d832f58f6f67417995c5fc725ff07e1459f52e"></a>

## Direct properties — gcp_cred_file.credential_file.blindfold_secret_info / cf8bec656890 / 3

<a id="canonical-6289febd13613ab53d388708b1bc1372a00d19f4455f90e8df3a69b6f53ec92e"></a>

<a id="canonical-dbc5a8f3e410a1ed47dcaa81f7294b151184baf8f78369fb7442b06120a2a92b"></a>

## decryption_provider property — gcp_cred_file.credential_file.blindfold_secret_info / cf8bec656890 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f64bb346d40676d9db79da34f6edad0ba90c4d05950a4f465cfd9f1fcdd21e9a"></a>

<a id="canonical-99d772d3b790fefd8a151d53ae2b104e23b049b8e6f40ebb1473a3ae6dbc9635"></a>

## location property — gcp_cred_file.credential_file.blindfold_secret_info / cf8bec656890 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-b8ed2d6cfdc48c146b10dd2cd1e09ed98ffb4cba7d7516ee99c3ec6ece69ac71"></a>

<a id="canonical-39ab962f466ee3a7bc76971e769577ce53fc33ee754d8ec8d28cab611af5ac73"></a>

## store_provider property — gcp_cred_file.credential_file.blindfold_secret_info / cf8bec656890 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9bf1243d3494d17a23b3a51e9ddde6c288679c45e0394fa1a6b687310da8d8a8"></a>

## Next pages — gcp_cred_file.credential_file.blindfold_secret_info / cf8bec656890 / 7

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)

<a id="canonical-ff665e06caa8b4391b73df4d1c574f32cb86ff5d2fc474f35611593749454c1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dede5e7aaa5925ed1ce2fde2d934a1862b40f76b11cd4432ed0222f3eef58a9"></a>

## gcp_cred_file.credential_file.clear_secret_info — gcp_cred_file.credential_file.clear_secret_info / 195c82229fc2 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-6a50e675849f84a504e89f06cea7867e15bac9423afe32c156fc99ff33548f03)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-8dfc6a8910f99bbfb26ba25097dca75e6f72f8ff1abd9e94b38131ad631c9ce8)
- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221)
- gcp_cred_file.credential_file.clear_secret_info

<a id="canonical-81f180237e2234202289bba78b63887dfbb71fb52d5b7de5b5c336b81c78bff4"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7abddf254db1357f830a981dcf5387292e5dd597e2ce505a2b2170b9aa3ec0af"></a>

## Direct properties — gcp_cred_file.credential_file.clear_secret_info / 195c82229fc2 / 3

<a id="canonical-07c7d7c7052c3e7058d871ed61c45abd550b235e5f3cc1d1b637a119ce813e3f"></a>

<a id="canonical-e4fe223c02f18e59aa9357389b3cbdc390de4d4c14cd13c61a19f3d6943af7e3"></a>

## provider_ref property — gcp_cred_file.credential_file.clear_secret_info / 195c82229fc2 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4236672525f9008e075c97b073b862b9a9609ef5e5ab7f3fadc3c4f7ba136114"></a>

<a id="canonical-0865aa82c3a5f3972ac34ab51c6f938978ef71496f22325d39fcd8a0f84a325e"></a>

## url property — gcp_cred_file.credential_file.clear_secret_info / 195c82229fc2 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-4a13bf0b6994160c2a6023b667271e09a3aa9cba6642beb2ddca57c6135291cf"></a>

## Next pages — gcp_cred_file.credential_file.clear_secret_info / 195c82229fc2 / 6

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-a4e2802dfd3404f5fc85820ef2db56c860b3417b9ee3bb532ddf808fb18c5221)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-947adf05823af193f5f3ffd1e35f5c1c03bff69b37a653a502bb59c1851d75f7)
