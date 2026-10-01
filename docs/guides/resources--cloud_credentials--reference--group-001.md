---
page_title: "xcsh_cloud_credentials reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials reference."
---

# xcsh_cloud_credentials reference

<a id="canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6607bbfb458dced3fcf4196a013bb4a14b59f42bad8f3ace773a4361ad38cfbe"></a>

## Property reference — Property reference / 959e61ea3a47 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- Property reference

<a id="canonical-98c0e98e38ff2989a10a4daf9a1495ff4df7ea08fa72708a364335ae5cfed1e4"></a>

## Direct properties — Property reference / 959e61ea3a47 / 3

<a id="canonical-86800da7becaf8bff327b925a391101989626158591a266711e061d2fc4acc77"></a>

<a id="canonical-5315aadcba035deffec9f22c7f651aa7c4b2d7298c83310c55559ca35ce64c05"></a>

## annotations property — Property reference / 959e61ea3a47 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53): complete subsection reference.

- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65): complete subsection reference.

- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b): complete subsection reference.

- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4): complete subsection reference.

<a id="canonical-ae655018076e6d2eba7c77fd372f4cb46f0f62db73f6ba87287a1d9d4f43af7d"></a>

<a id="canonical-653fe805b85fe2276c78c89e935e72b900d599148eb69962cf11c99fb85dd923"></a>

## description property — Property reference / 959e61ea3a47 / 5

Type: `"string"`. Optional.

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

<a id="canonical-23c4c5880c09d2b45ed37791f364d90af1a0a8ab6abb7648880725fb58ebee43"></a>

<a id="canonical-476972e5a134cccfc33dc930daf78ea2b32461bcbcdd2411994a4cd3fb9bf521"></a>

## disable property — Property reference / 959e61ea3a47 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7): complete subsection reference.

<a id="canonical-64e7469fcf37417e04a4c094512069ec6c4a0de66929057cb6f6f4276100ff1c"></a>

<a id="canonical-323ed11c045e3ee2936d85248134f6d816c56829e49319a55492b22749ba7f83"></a>

## id property — Property reference / 959e61ea3a47 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-208db0e17f02ba45451c3b6c9cc0f339b20ba70eee1f461d246c4940ef03fcf3"></a>

<a id="canonical-a109c38030ed9a7dbd27883fb19fded6b7f87fbd7b8e420b3d9b5110790a80ae"></a>

## labels property — Property reference / 959e61ea3a47 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-dd522debe24c946c486b84d5c3012f47ae9ae2044d28b5ceb54229789c9983bb"></a>

<a id="canonical-ca43b599a35678cb0e27caaf4f93cdfe6d1a6b121ef87cde6a094f145163ad9f"></a>

## name property — Property reference / 959e61ea3a47 / 9

Type: `"string"`. Required.

Name of the Cloud Credentials. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-b35ec70551a16b082e7da2a76adca3d66eb7715dd7e9fad5093b19c0ec6a7368"></a>

<a id="canonical-39e8873da6cc463e372e49c9e6dcef08ef2203f77f6db29c7f4fbf5ab97fecd9"></a>

## namespace property — Property reference / 959e61ea3a47 / 10

Type: `"string"`. Required.

Namespace where the Cloud Credentials is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--cloud_credentials--reference--group-001.md#canonical-3acc6c96cd5416a3e10b6797715d00081c043429c9034ec56cfd26dcf256fc19): complete subsection reference.

<a id="canonical-818fb5a3e1faf1c549727b3385d70829353b30284c4b05fe1bc26f44dfb8881f"></a>

## All schema paths — Property reference / 959e61ea3a47 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_credentials--reference--group-001.md#canonical-86800da7becaf8bff327b925a391101989626158591a266711e061d2fc4acc77) |
| `aws_assume_role` | [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-84f22769fd3e174d0fe7c64ffde42b0a1cd12ff841df03686fbad29d20fd5132) |
| `aws_assume_role.custom_external_id` | [aws_assume_role.custom_external_id](resources--cloud_credentials--reference--group-001.md#canonical-12b70d4d1a852e86a48de6d4774b1cfccf463f54659ad723539165475989eac7) |
| `aws_assume_role.duration_seconds` | [aws_assume_role.duration_seconds](resources--cloud_credentials--reference--group-001.md#canonical-e02fc8b61732cd346e409f6dc9e23455943643252c3158c11421c779a12eb8ca) |
| `aws_assume_role.external_id_is_optional` | [aws_assume_role.external_id_is_optional](resources--cloud_credentials--reference--group-001.md#canonical-74186f60bb062d070a7029ece4687c108670f5d34a5c3c58d05ec10efdb462b6) |
| `aws_assume_role.external_id_is_tenant_id` | [aws_assume_role.external_id_is_tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-e14e5510df168cf3714396e8fc43e56a4f46481362b177773b775482c7ed1bad) |
| `aws_assume_role.role_arn` | [aws_assume_role.role_arn](resources--cloud_credentials--reference--group-001.md#canonical-26de7af59bfb04fcde52246383c2e4f8bc81f16ade3494410eb4b45bbbaae2e7) |
| `aws_assume_role.session_name` | [aws_assume_role.session_name](resources--cloud_credentials--reference--group-001.md#canonical-88a400c4296ce616be729d9967897d69ddf8ee52b7c633e365283ab9f6924def) |
| `aws_assume_role.session_tags` | [aws_assume_role.session_tags](resources--cloud_credentials--reference--group-001.md#canonical-2f419eacb0cbadb123511a63d487d5209e4c55cb68d22cb0684e0e20ac6075f6) |
| `aws_secret_key` | [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0b2eb6ef14be44880629a24347d679384ecd12b66af698aeb7e40711f5cba768) |
| `aws_secret_key.access_key` | [aws_secret_key.access_key](resources--cloud_credentials--reference--group-001.md#canonical-acb081a5fd640910c9853e14ee84eae1edcf2c1767a8d797ff3259d9e78b8031) |
| `aws_secret_key.secret_key` | [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-59ff61eae9655c0d66f9e12d9db9ac4119ec1729245846cffc58a105c7038ab6) |
| `aws_secret_key.secret_key.blindfold_secret_info` | [aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-d74254856bf40b3160c9cedf0e83f251cc2e7bfd878a4f31f57be30c57717e63) |
| `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-cd5ccc8216c671fa23b9d5b72c700cb8298299850b52ccc5d07f84355d074f68) |
| `aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_secret_key.secret_key.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-b254d99e0db281db9eec1611f82d3a325f7543b15f8a2fcceb65928cb0d8bca9) |
| `aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_secret_key.secret_key.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-86e49df6df043a920ec36e074b907e8409d25844183821f3f530e5d1de52d61b) |
| `aws_secret_key.secret_key.clear_secret_info` | [aws_secret_key.secret_key.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-55c6f99009883f9d4e709478087460315ab0ab7302e203ef44ec197d6bed4b31) |
| `aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_secret_key.secret_key.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-4ea92a53177db402f5a58b134e2b0d6564a4843c5e10686a8464380f6e750d76) |
| `aws_secret_key.secret_key.clear_secret_info.url` | [aws_secret_key.secret_key.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-110d54cad3f7d789b4995b960bd4f421d95dbdeb5086b7d7ee346a4068b0115f) |
| `azure_client_secret` | [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-f4fcb5013edfbbd2e1f929d412f207c8f737feb5fe11334499701c98828f5296) |
| `azure_client_secret.client_id` | [azure_client_secret.client_id](resources--cloud_credentials--reference--group-001.md#canonical-ccc7c8b67d1e2b868167a4cd5c4302fda59b16e0d160634e15252a183452d142) |
| `azure_client_secret.client_secret` | [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-124d3f1223fdfba6a17f2a20cc0f8c9eaca0f8f6c51c7dbb960214f2550566ef) |
| `azure_client_secret.client_secret.blindfold_secret_info` | [azure_client_secret.client_secret.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-f1ea43faa12bcfec73ed945790136a74b881045a5e556adcda743accc01b2f13) |
| `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` | [azure_client_secret.client_secret.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-5304bc1622911374fa98c9a357a9bdc057f656a8a74052bc80508ac7ccd7545e) |
| `azure_client_secret.client_secret.blindfold_secret_info.location` | [azure_client_secret.client_secret.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-024680e4e65413e5b59052710cd1014feb4f8ee9da53372350875cdabe62f1ae) |
| `azure_client_secret.client_secret.blindfold_secret_info.store_provider` | [azure_client_secret.client_secret.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-e1f2c59d41fc6cf3ce68b535cd40ef822cf762c33f32ad050f13046226e7c8e9) |
| `azure_client_secret.client_secret.clear_secret_info` | [azure_client_secret.client_secret.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-fd49bc736890090e5311a4dc776b3144142abb96f214258a5ecdab6f4eb651c7) |
| `azure_client_secret.client_secret.clear_secret_info.provider_ref` | [azure_client_secret.client_secret.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-97e37a9df31f2458c0fbc51559ad609c59e6283dd6367718064c791900ee928b) |
| `azure_client_secret.client_secret.clear_secret_info.url` | [azure_client_secret.client_secret.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-a888d1b281e51b43274bf4d4aa1a27cd6d4a037955b3cdfb6479d05be37d161e) |
| `azure_client_secret.subscription_id` | [azure_client_secret.subscription_id](resources--cloud_credentials--reference--group-001.md#canonical-8fb62b6c8297ebb01aaf86dcc1b811c53aef8a4ccb5525acd8859d5dc9b43303) |
| `azure_client_secret.tenant_id` | [azure_client_secret.tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-397ac8ee04bc2d26508fc35daaa0f6fe1858c0a63ba0b582ac7bf7c5d3d6aacc) |
| `azure_pfx_certificate` | [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-df2c0182d927c3ce0cc417317dff1a750d40269b6c6bb70446a50d29b8446755) |
| `azure_pfx_certificate.certificate_url` | [azure_pfx_certificate.certificate_url](resources--cloud_credentials--reference--group-001.md#canonical-617e646d802c8b308d378063a389c7fb4392a3bc618bc78a4b67a2e35bb55f62) |
| `azure_pfx_certificate.client_id` | [azure_pfx_certificate.client_id](resources--cloud_credentials--reference--group-001.md#canonical-efb91a6b55c67a03637d8fd34ee14fdee5bde589203423c4ea046d93dd37432c) |
| `azure_pfx_certificate.password` | [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-6012d8121a30b2c4037ae600f862ffef601d41ba9b1b32adbf87200d4fe4dabd) |
| `azure_pfx_certificate.password.blindfold_secret_info` | [azure_pfx_certificate.password.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-dea3b516b0147de14036ccd2739f1a14f0ef9b4cf574909ff38615a613fd7502) |
| `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` | [azure_pfx_certificate.password.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-3008d1320b299b99b3ba36d39ad1dfad2f878dd966fb3dcf425f243c2b68e0aa) |
| `azure_pfx_certificate.password.blindfold_secret_info.location` | [azure_pfx_certificate.password.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-144511479c49e621748778433d201504bd37ab0a61d9d1a1ecf76cca13e7d28e) |
| `azure_pfx_certificate.password.blindfold_secret_info.store_provider` | [azure_pfx_certificate.password.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-cfe18e7a64b37e5b8a5bb3c0c33706821d8ae62e48333e611117efc1d2f4b8b9) |
| `azure_pfx_certificate.password.clear_secret_info` | [azure_pfx_certificate.password.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-8ea8654b2db3f2bac34109e7b363b31db9ecd1cd0f02c6fac79a34a33c3197be) |
| `azure_pfx_certificate.password.clear_secret_info.provider_ref` | [azure_pfx_certificate.password.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-93fde45572ba704d7c341713d8b54842333bd3871bfc5139f54982f916e8e7ef) |
| `azure_pfx_certificate.password.clear_secret_info.url` | [azure_pfx_certificate.password.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-7f868bf8f38d544528f4423a4c55b56689dc8ad466a13174e5ed4e3020a19409) |
| `azure_pfx_certificate.subscription_id` | [azure_pfx_certificate.subscription_id](resources--cloud_credentials--reference--group-001.md#canonical-565ccbb81151ecc92ba38a2b098fcf3666c422177fa5eb26caff5eaff520ac6f) |
| `azure_pfx_certificate.tenant_id` | [azure_pfx_certificate.tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-b5e173f95b735e65cc012dc756a06ab729ec17969362589cc2d3a552caa4ec29) |
| `description` | [description](resources--cloud_credentials--reference--group-001.md#canonical-ae655018076e6d2eba7c77fd372f4cb46f0f62db73f6ba87287a1d9d4f43af7d) |
| `disable` | [disable](resources--cloud_credentials--reference--group-001.md#canonical-23c4c5880c09d2b45ed37791f364d90af1a0a8ab6abb7648880725fb58ebee43) |
| `gcp_cred_file` | [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-c6d5e0047cb0140207434ad4cacf0c414ded9691a9e46f5ae652d6d038bd5e86) |
| `gcp_cred_file.credential_file` | [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-112ae79be7b74f472b49b8bf1962c588330c1d56f688dcc30a796c8683cda5b7) |
| `gcp_cred_file.credential_file.blindfold_secret_info` | [gcp_cred_file.credential_file.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-b73371e9932f7968ba9b9f87dce78fcc42b4749cf0f091060650c1207f142984) |
| `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-ded1bde5f500a00b68faf3964668681561389ac608cfd29333883e69de64e9fe) |
| `gcp_cred_file.credential_file.blindfold_secret_info.location` | [gcp_cred_file.credential_file.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-0741dafc23007cfaaff70fa0a316f14f662e597f96d3b17cfb57b67b007f4020) |
| `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-dfa60a07e3198c82c322cf47b2ec05614caa1a7c9377bad249bea41c1d4909e4) |
| `gcp_cred_file.credential_file.clear_secret_info` | [gcp_cred_file.credential_file.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-c9b120cbfd1d032cfe4cf254f8b6ca24eed2f503d1aa499a0903aa3a4a4fc75e) |
| `gcp_cred_file.credential_file.clear_secret_info.provider_ref` | [gcp_cred_file.credential_file.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-76cb4f180384af9f43aca561452bbc983599de9fdff772d921458d0dfabc0c5d) |
| `gcp_cred_file.credential_file.clear_secret_info.url` | [gcp_cred_file.credential_file.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-139f8e150e8996473fa4c6a9c13af0fc64d971830e3bdec4290470ae96724926) |
| `id` | [id](resources--cloud_credentials--reference--group-001.md#canonical-64e7469fcf37417e04a4c094512069ec6c4a0de66929057cb6f6f4276100ff1c) |
| `labels` | [labels](resources--cloud_credentials--reference--group-001.md#canonical-208db0e17f02ba45451c3b6c9cc0f339b20ba70eee1f461d246c4940ef03fcf3) |
| `name` | [name](resources--cloud_credentials--reference--group-001.md#canonical-dd522debe24c946c486b84d5c3012f47ae9ae2044d28b5ceb54229789c9983bb) |
| `namespace` | [namespace](resources--cloud_credentials--reference--group-001.md#canonical-b35ec70551a16b082e7da2a76adca3d66eb7715dd7e9fad5093b19c0ec6a7368) |
| `timeouts` | [timeouts](resources--cloud_credentials--reference--group-001.md#canonical-44c9716c36756ca2eace2ede36b56d0fe85dd0d549e848b75bb5a838c5abf87f) |
| `timeouts.create` | [timeouts.create](resources--cloud_credentials--reference--group-001.md#canonical-dbf5f1b102f11a69bce280da5449e40f442c2ad59f27b9b34127e96df1aa73d3) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_credentials--reference--group-001.md#canonical-8c4eb1fe1f24394877fd2beed9dfc1131d5bbdfd2263553890eead22ca0f82d7) |
| `timeouts.read` | [timeouts.read](resources--cloud_credentials--reference--group-001.md#canonical-d48b89934793851663916b540810463da14d022e42ed89cfd3d6bd185ca05cb1) |
| `timeouts.update` | [timeouts.update](resources--cloud_credentials--reference--group-001.md#canonical-de2c35d1fbb4fb4407a978591fefcf8c73defbaffa1ccd59191405871e2ad3ae) |

<a id="canonical-ff16e97103a181c6224554e0c85a59b8bfefc13c43352c5e53a31a20757cb357"></a>

## Next pages — Property reference / 959e61ea3a47 / 12

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7)
- [timeouts](resources--cloud_credentials--reference--group-001.md#canonical-3acc6c96cd5416a3e10b6797715d00081c043429c9034ec56cfd26dcf256fc19)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b59c2da63724bdbf247adeb2e319369d161f004d39c2f623b9a89be1eeacb84"></a>

## aws_assume_role — aws_assume_role / 36963b08a4c2 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- aws_assume_role

<a id="canonical-84f22769fd3e174d0fe7c64ffde42b0a1cd12ff841df03686fbad29d20fd5132"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_assume\_role, aws\_secret\_key, azure\_client\_secret, azure\_pfx\_certificate,
gcp\_cred\_file\] AWS Assume Role to Handle Delegated Access.

Upstream description:

AWS Assume Role to Handle Delegated Access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration_seconds",
    "role_arn",
    "session_name"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_optional"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_tenant_id"),
  validators.ConflictingObjectAttributes("external_id_is_optional",
    "external_id_is_tenant_id")}
```

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

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-84f22769fd3e174d0fe7c64ffde42b0a1cd12ff841df03686fbad29d20fd5132)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0b2eb6ef14be44880629a24347d679384ecd12b66af698aeb7e40711f5cba768)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-f4fcb5013edfbbd2e1f929d412f207c8f737feb5fe11334499701c98828f5296)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-df2c0182d927c3ce0cc417317dff1a750d40269b6c6bb70446a50d29b8446755)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-c6d5e0047cb0140207434ad4cacf0c414ded9691a9e46f5ae652d6d038bd5e86)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_assume_role {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6030bee2c7036e46598bad456e1637830a4fe23418b96127a971ba39d9c3040"></a>

## Direct properties — aws_assume_role / 36963b08a4c2 / 3

<a id="canonical-12b70d4d1a852e86a48de6d4774b1cfccf463f54659ad723539165475989eac7"></a>

<a id="canonical-2d16d61f7ffba85c716a6d9ac63f9274717b29690301db3ca3fc481553b3990b"></a>

## custom_external_id property — aws_assume_role / 36963b08a4c2 / 4

Type: `"string"`. Optional.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

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

<a id="canonical-e02fc8b61732cd346e409f6dc9e23455943643252c3158c11421c779a12eb8ca"></a>

<a id="canonical-ae5db9010f3878f5ab115e0b3b7a4e2b4034067ed44045c06326540697fc7d09"></a>

## duration_seconds property — aws_assume_role / 36963b08a4c2 / 5

Type: `"number"`. Optional.

The duration, in seconds of the role session.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(3600, 43200),
}
```

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

- [external_id_is_optional](resources--cloud_credentials--reference--group-001.md#canonical-dab3eb68bb7cfc99900c5b903d0b44a8a3702ed5d067879517b62ceaa87ebb64): complete subsection reference.

- [external_id_is_tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-a34e5cf0ce0ea91d1d99a1ec37cf38534bf58bff77a2163b4ef3376b03fb0576): complete subsection reference.

<a id="canonical-26de7af59bfb04fcde52246383c2e4f8bc81f16ade3494410eb4b45bbbaae2e7"></a>

<a id="canonical-080d4384e6a000ba32dfaace65993c7ce8cfd298a59374de1295ba13eb61b3c8"></a>

## role_arn property — aws_assume_role / 36963b08a4c2 / 6

Type: `"string"`. Optional.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 2048),
}
```

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

<a id="canonical-88a400c4296ce616be729d9967897d69ddf8ee52b7c633e365283ab9f6924def"></a>

<a id="canonical-55da906b74c42ebe54f59b68ba9b748c2f5b45edfce599e25bc84958bb9adffd"></a>

## session_name property — aws_assume_role / 36963b08a4c2 / 7

Type: `"string"`. Optional.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

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

<a id="canonical-2f419eacb0cbadb123511a63d487d5209e4c55cb68d22cb0684e0e20ac6075f6"></a>

<a id="canonical-93510d15ccd73feb09e52d6f03054e148853aa5e5b512483967488069964c811"></a>

## session_tags property — aws_assume_role / 36963b08a4c2 / 8

Type: `["map", "string"]`. Optional.

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

<a id="canonical-f8dabc2301a6da0e2ef6d846e6a65a1490234e714debff84eca868614aacb89b"></a>

## Next pages — aws_assume_role / 36963b08a4c2 / 9

- [aws_assume_role.external_id_is_optional](resources--cloud_credentials--reference--group-001.md#canonical-dab3eb68bb7cfc99900c5b903d0b44a8a3702ed5d067879517b62ceaa87ebb64)
- [aws_assume_role.external_id_is_tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-a34e5cf0ce0ea91d1d99a1ec37cf38534bf58bff77a2163b4ef3376b03fb0576)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-dab3eb68bb7cfc99900c5b903d0b44a8a3702ed5d067879517b62ceaa87ebb64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af85574bd410f8c6ca2b02b40146ba4e5e1e619cc28e3eb1d003232c328c19a3"></a>

## aws_assume_role.external_id_is_optional — aws_assume_role.external_id_is_optional / e2e8228a1d18 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53)
- aws_assume_role.external_id_is_optional

<a id="canonical-74186f60bb062d070a7029ece4687c108670f5d34a5c3c58d05ec10efdb462b6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
external_id_is_optional = {}
```

<a id="canonical-b45ac8b7755a7e94306912c863dbadc32ee51a64be6e6de8e2d35f70298e3c5d"></a>

## Direct properties — aws_assume_role.external_id_is_optional / e2e8228a1d18 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-77e61fd39b57b009fd1153531ed610ab5ff3cfeae0ddf3fb2471eaeb7b88f371"></a>

## Next pages — aws_assume_role.external_id_is_optional / e2e8228a1d18 / 4

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-a34e5cf0ce0ea91d1d99a1ec37cf38534bf58bff77a2163b4ef3376b03fb0576"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6b81769a8dfe8bf381397480181dae1caaf2510017529a1157d03d32b751113"></a>

## aws_assume_role.external_id_is_tenant_id — aws_assume_role.external_id_is_tenant_id / 62e70a5438ad / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53)
- aws_assume_role.external_id_is_tenant_id

<a id="canonical-e14e5510df168cf3714396e8fc43e56a4f46481362b177773b775482c7ed1bad"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
external_id_is_tenant_id = {}
```

<a id="canonical-4063f9a376775feb30a91b4059e249906f204ddba8bb0afff36098ac924988a7"></a>

## Direct properties — aws_assume_role.external_id_is_tenant_id / 62e70a5438ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d5c8a3220d2cb023112348dfbf234c33ede57cb196fb8a87f854232f534b30c"></a>

## Next pages — aws_assume_role.external_id_is_tenant_id / 62e70a5438ad / 4

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-e1774c8b1a1466c11c75cb97ee7d09fcc6af606dc9d9ae38f719ba05cfd63d53)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-119b320ea4f0cfb9d14a3376e71a6fda2e24e0768021f309f6cc33513449bb02"></a>

## aws_secret_key — aws_secret_key / 4cb92e93af1d / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- aws_secret_key

<a id="canonical-0b2eb6ef14be44880629a24347d679384ecd12b66af698aeb7e40711f5cba768"></a>

Type: `"object"`. single nested block, Optional.

AWS Programmatic Access Credentials type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("access_key")}
```

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

Terraform syntax:

```terraform
aws_secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3870b1833e8da9da4dfbf1fa5fdeeb92a2818033cacdd20821376900059703bd"></a>

## Direct properties — aws_secret_key / 4cb92e93af1d / 3

<a id="canonical-acb081a5fd640910c9853e14ee84eae1edcf2c1767a8d797ff3259d9e78b8031"></a>

<a id="canonical-d3b725f8d97b118470a0b3fa194232625151185e0ed05fc451c648b4da316143"></a>

## access_key property — aws_secret_key / 4cb92e93af1d / 4

Type: `"string"`. Optional.

Access Key ID. Access key ID for your AWS account.

Upstream description:

Access key ID for your AWS account.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [secret_key](resources--cloud_credentials--reference--group-001.md#canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43): complete subsection reference.

<a id="canonical-14310ca05713a1346312ff2342748c7b22588b174a3c365a2a3c01e86e7ab1ad"></a>

## Next pages — aws_secret_key / 4cb92e93af1d / 5

- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d33d48f1e5c2aa948f60dbd0731f02a222236c5478dc2113b55720d8e02bde44"></a>

## aws_secret_key.secret_key — aws_secret_key.secret_key / de566b3b2368 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65)
- aws_secret_key.secret_key

<a id="canonical-59ff61eae9655c0d66f9e12d9db9ac4119ec1729245846cffc58a105c7038ab6"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-084fc15628331d12474b9557536ea272028910f19143e39aa30f56afa7eeaed7"></a>

## Direct properties — aws_secret_key.secret_key / de566b3b2368 / 3

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-ff85862deb431f501bef683ce3eb785ceeab2389501713a4367d6d0ab9b41137): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-0638ebae88296b2a3f71b61e30472cacd6569174219c51c622697335d3ea2ac9): complete subsection reference.

<a id="canonical-1854a24916db7a8d4580225317d604ed4ad67aae1ec2fd76c68c42a1c766ed55"></a>

## Next pages — aws_secret_key.secret_key / de566b3b2368 / 4

- [aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-ff85862deb431f501bef683ce3eb785ceeab2389501713a4367d6d0ab9b41137)
- [aws_secret_key.secret_key.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-0638ebae88296b2a3f71b61e30472cacd6569174219c51c622697335d3ea2ac9)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-ff85862deb431f501bef683ce3eb785ceeab2389501713a4367d6d0ab9b41137"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ceb08fd9f01be1c73027581a4ef577a98efa35a4be9333befc2dc714858fb09"></a>

## aws_secret_key.secret_key.blindfold_secret_info — aws_secret_key.secret_key.blindfold_secret_info / 7bcf0577d9fd / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65)
- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43)
- aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-d74254856bf40b3160c9cedf0e83f251cc2e7bfd878a4f31f57be30c57717e63"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-d92a4b88272f04e7baaaa266c6413e28f48bb7fc5eb5263258bad24d09125699"></a>

## Direct properties — aws_secret_key.secret_key.blindfold_secret_info / 7bcf0577d9fd / 3

<a id="canonical-cd5ccc8216c671fa23b9d5b72c700cb8298299850b52ccc5d07f84355d074f68"></a>

<a id="canonical-b5e1c3c2032150e31fa2a49a5faebc51a79a8462454d5b7822b3882bce8cd436"></a>

## decryption_provider property — aws_secret_key.secret_key.blindfold_secret_info / 7bcf0577d9fd / 4

Type: `"string"`. Optional.

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

<a id="canonical-b254d99e0db281db9eec1611f82d3a325f7543b15f8a2fcceb65928cb0d8bca9"></a>

<a id="canonical-da75c6c03c0b051f19f6dffcf2c033586ec34aaffdef28eb90a9e880e09e3d30"></a>

## location property — aws_secret_key.secret_key.blindfold_secret_info / 7bcf0577d9fd / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-86e49df6df043a920ec36e074b907e8409d25844183821f3f530e5d1de52d61b"></a>

<a id="canonical-ee560f3e81e7ed513ff1dae4e145e4fdefa9daefc1a5a96d31900c4b8e11749b"></a>

## store_provider property — aws_secret_key.secret_key.blindfold_secret_info / 7bcf0577d9fd / 6

Type: `"string"`. Optional.

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

<a id="canonical-ed1c9a36b5f7ac0e4bae10b8659e981c489d150ed8d6b93eca9dfc3798b7c058"></a>

## Next pages — aws_secret_key.secret_key.blindfold_secret_info / 7bcf0577d9fd / 7

- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-0638ebae88296b2a3f71b61e30472cacd6569174219c51c622697335d3ea2ac9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6aa93430a1582e834013b2d7cccb5368bc7f386f1c57b604f386cac8eec83d4e"></a>

## aws_secret_key.secret_key.clear_secret_info — aws_secret_key.secret_key.clear_secret_info / 7bde656c5413 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-ef5772db96062a5b437953597113fc2ea170e7f7e18eaeac2c850abb9b133a65)
- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43)
- aws_secret_key.secret_key.clear_secret_info

<a id="canonical-55c6f99009883f9d4e709478087460315ab0ab7302e203ef44ec197d6bed4b31"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-64b5ce03a046ead4ac5c144a4ed988530ddc017e40d96c9f1958b07c982d9925"></a>

## Direct properties — aws_secret_key.secret_key.clear_secret_info / 7bde656c5413 / 3

<a id="canonical-4ea92a53177db402f5a58b134e2b0d6564a4843c5e10686a8464380f6e750d76"></a>

<a id="canonical-2414f71e005e3f797cdfba4c3738687fd6e3cbc333afffa682115f68cb694e90"></a>

## provider_ref property — aws_secret_key.secret_key.clear_secret_info / 7bde656c5413 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-110d54cad3f7d789b4995b960bd4f421d95dbdeb5086b7d7ee346a4068b0115f"></a>

<a id="canonical-3e37f3640234eb37d5dd307fa190d0a24675a0dc39466aff16ce73a12807c06c"></a>

## url property — aws_secret_key.secret_key.clear_secret_info / 7bde656c5413 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-dfa1e19c223cf99ac229a79dcb8c7c5513714551fdc8119ddb1ea693582f7edb"></a>

## Next pages — aws_secret_key.secret_key.clear_secret_info / 7bde656c5413 / 6

- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-277549eb2bb912fb9bd5f951e1e1573049298027d498a8fdb711d0eb82249f43)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41bd3cf38fee4234bbd134636af2793802b6f41e34286c3388007f9c3daf8bd7"></a>

## azure_client_secret — azure_client_secret / f59086c5f30f / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- azure_client_secret

<a id="canonical-f4fcb5013edfbbd2e1f929d412f207c8f737feb5fe11334499701c98828f5296"></a>

Type: `"object"`. single nested block, Optional.

Azure Client Secret. Azure Credentials Client Secret type.

Upstream description:

Azure Credentials Client Secret type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("client_id",
    "subscription_id",
    "tenant_id")}
```

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

Terraform syntax:

```terraform
azure_client_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6435fbdceba4edcd7cf2db77b597285a97ebf12decaff67e27d98f09318a491"></a>

## Direct properties — azure_client_secret / f59086c5f30f / 3

<a id="canonical-ccc7c8b67d1e2b868167a4cd5c4302fda59b16e0d160634e15252a183452d142"></a>

<a id="canonical-938bc5e1fd2eeda07c3f77ff3db0845072640a8ecfed94250b5f69cfc92a1f96"></a>

## client_id property — azure_client_secret / f59086c5f30f / 4

Type: `"string"`. Optional.

Client ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [client_secret](resources--cloud_credentials--reference--group-001.md#canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960): complete subsection reference.

<a id="canonical-8fb62b6c8297ebb01aaf86dcc1b811c53aef8a4ccb5525acd8859d5dc9b43303"></a>

<a id="canonical-a717763ddeba678fd2455c680bf1c22a271cb728e0c36efce2cd536e57c39ede"></a>

## subscription_id property — azure_client_secret / f59086c5f30f / 5

Type: `"string"`. Optional.

Subscription ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-397ac8ee04bc2d26508fc35daaa0f6fe1858c0a63ba0b582ac7bf7c5d3d6aacc"></a>

<a id="canonical-e172a44b9e3b6bd0c26ab3a9a965f301b3bbf34fc6acf4423076ba87e7dead5b"></a>

## tenant_id property — azure_client_secret / f59086c5f30f / 6

Type: `"string"`. Optional.

Tenant ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-ee45b2e9500077b68603233f691f594d01b53ee0aaa58b1273359e5866551685"></a>

## Next pages — azure_client_secret / f59086c5f30f / 7

- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07e265e31258d75bc43c19a4c3c5f415db94b9f4a12de22a877be7c291d08f0c"></a>

## azure_client_secret.client_secret — azure_client_secret.client_secret / 9b70490ffc9b / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b)
- azure_client_secret.client_secret

<a id="canonical-124d3f1223fdfba6a17f2a20cc0f8c9eaca0f8f6c51c7dbb960214f2550566ef"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
client_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-fdcf50bde4983ffd8e1b01f46d77336b36f7aadd60a5bb18d32c84d81f9d4300"></a>

## Direct properties — azure_client_secret.client_secret / 9b70490ffc9b / 3

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-1c84d2a4251eac503fa95bd879f427500f3d3c9ad7c866ea9f8172c952ff127c): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-57072146da50875d392c75350bfdb5deba2f8251bd23db9800f4c0c06185d765): complete subsection reference.

<a id="canonical-8f9cd20e1437e85411e2110138e2b0e6471140af7ca3bdacfad1c439b1c7b02b"></a>

## Next pages — azure_client_secret.client_secret / 9b70490ffc9b / 4

- [azure_client_secret.client_secret.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-1c84d2a4251eac503fa95bd879f427500f3d3c9ad7c866ea9f8172c952ff127c)
- [azure_client_secret.client_secret.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-57072146da50875d392c75350bfdb5deba2f8251bd23db9800f4c0c06185d765)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-1c84d2a4251eac503fa95bd879f427500f3d3c9ad7c866ea9f8172c952ff127c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a62ce13be77114c2ff14bd80838315209ba5979c18f63c23c4f167bd33a8df6c"></a>

## azure_client_secret.client_secret.blindfold_secret_info — azure_client_secret.client_secret.blindfold_secret_info / 75b641714140 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b)
- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960)
- azure_client_secret.client_secret.blindfold_secret_info

<a id="canonical-f1ea43faa12bcfec73ed945790136a74b881045a5e556adcda743accc01b2f13"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-10b81aecbcd7787f98601eb249922e6e26cd9e4c6a9017e6ce38a951ce094936"></a>

## Direct properties — azure_client_secret.client_secret.blindfold_secret_info / 75b641714140 / 3

<a id="canonical-5304bc1622911374fa98c9a357a9bdc057f656a8a74052bc80508ac7ccd7545e"></a>

<a id="canonical-5d365551696a089464673a5e5e2c4b7197515a7961e9ba90a90ff97cdc11131b"></a>

## decryption_provider property — azure_client_secret.client_secret.blindfold_secret_info / 75b641714140 / 4

Type: `"string"`. Optional.

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

<a id="canonical-024680e4e65413e5b59052710cd1014feb4f8ee9da53372350875cdabe62f1ae"></a>

<a id="canonical-0fc0f0f726dc073ad81b581a3c0d75623884b497cc1be56d2b35b28860662b27"></a>

## location property — azure_client_secret.client_secret.blindfold_secret_info / 75b641714140 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-e1f2c59d41fc6cf3ce68b535cd40ef822cf762c33f32ad050f13046226e7c8e9"></a>

<a id="canonical-1edb2988cdc3e4453cd8cac2f507aa5ab194a4b7e215a96bd9f090b4a47b3df3"></a>

## store_provider property — azure_client_secret.client_secret.blindfold_secret_info / 75b641714140 / 6

Type: `"string"`. Optional.

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

<a id="canonical-7f88b2d97138669e36a3a900c5e193eaf897316dd6612244a1ac5191ea3b3420"></a>

## Next pages — azure_client_secret.client_secret.blindfold_secret_info / 75b641714140 / 7

- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-57072146da50875d392c75350bfdb5deba2f8251bd23db9800f4c0c06185d765"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ce6e554131eaf65ca5f1e9d2ee7df400381d752018c6d577382a9998653b756"></a>

## azure_client_secret.client_secret.clear_secret_info — azure_client_secret.client_secret.clear_secret_info / cf884f88bae1 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0d0ea7d2e8a2ccade98d84da9e05785e901de427a4e1bcabd8f887f5cd2fa27b)
- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960)
- azure_client_secret.client_secret.clear_secret_info

<a id="canonical-fd49bc736890090e5311a4dc776b3144142abb96f214258a5ecdab6f4eb651c7"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-29888287a71b228355b3f2a400b2158ca25572b5f4e155e379ce2bb416712560"></a>

## Direct properties — azure_client_secret.client_secret.clear_secret_info / cf884f88bae1 / 3

<a id="canonical-97e37a9df31f2458c0fbc51559ad609c59e6283dd6367718064c791900ee928b"></a>

<a id="canonical-4cf8201b186b07db5086146507db6724492444bfcac3710d834654063edd44f3"></a>

## provider_ref property — azure_client_secret.client_secret.clear_secret_info / cf884f88bae1 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a888d1b281e51b43274bf4d4aa1a27cd6d4a037955b3cdfb6479d05be37d161e"></a>

<a id="canonical-af6d283775260ba413a7b8afa2cc07e1ee7dde122cfb21db8a03a30b07b8858c"></a>

## url property — azure_client_secret.client_secret.clear_secret_info / cf884f88bae1 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-6f8e5c23c1cfec38a44c248cb87e10ab976cb56f89f9864032392357aab010a1"></a>

## Next pages — azure_client_secret.client_secret.clear_secret_info / cf884f88bae1 / 6

- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-844239be548eccca761ee30e353773f7f5d22307b853badbe8909d532340a960)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05aaf65a71de2a26e4b8ddbb4b4d8682e9b9d02e92a5d237fab7936c755ae30e"></a>

## azure_pfx_certificate — azure_pfx_certificate / bdd75999ebad / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- azure_pfx_certificate

<a id="canonical-df2c0182d927c3ce0cc417317dff1a750d40269b6c6bb70446a50d29b8446755"></a>

Type: `"object"`. single nested block, Optional.

Azure Credentials Client Certificate type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url",
    "client_id",
    "subscription_id",
    "tenant_id")}
```

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

Terraform syntax:

```terraform
azure_pfx_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-4bcefa7b2e2181eff26b036adef34705b31e7c98d214eef1cc607fadf52b3e3d"></a>

## Direct properties — azure_pfx_certificate / bdd75999ebad / 3

<a id="canonical-617e646d802c8b308d378063a389c7fb4392a3bc618bc78a4b67a2e35bb55f62"></a>

<a id="canonical-0288862b36d3cd7068009a258271c5c48bfb59ad601970b123411e4642619f25"></a>

## certificate_url property — azure_pfx_certificate / bdd75999ebad / 4

Type: `"string"`. Optional.

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Upstream description:

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

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

<a id="canonical-efb91a6b55c67a03637d8fd34ee14fdee5bde589203423c4ea046d93dd37432c"></a>

<a id="canonical-aeb474ce7da3c982b15955b690be2a3e78a31ea0ac4f107eec96a0529d5eabc5"></a>

## client_id property — azure_pfx_certificate / bdd75999ebad / 5

Type: `"string"`. Optional.

Client ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [password](resources--cloud_credentials--reference--group-001.md#canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d): complete subsection reference.

<a id="canonical-565ccbb81151ecc92ba38a2b098fcf3666c422177fa5eb26caff5eaff520ac6f"></a>

<a id="canonical-113ccf186d5845b71db8f21a88beb9b94994ce9247db77b2aa99c97b0f76637c"></a>

## subscription_id property — azure_pfx_certificate / bdd75999ebad / 6

Type: `"string"`. Optional.

Subscription ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-b5e173f95b735e65cc012dc756a06ab729ec17969362589cc2d3a552caa4ec29"></a>

<a id="canonical-3b7c7ef74561f774f185e15da08db79f984a35ff2730138867c0d21fa9926e3e"></a>

## tenant_id property — azure_pfx_certificate / bdd75999ebad / 7

Type: `"string"`. Optional.

Tenant ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-0308e4a45e170743186662ec4606cd8bd02d67c40592e60d529fdac43f001b26"></a>

## Next pages — azure_pfx_certificate / bdd75999ebad / 8

- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15fb8248c7e4e65db97887b3d01a949beaa784b069e99f55e9795c5744543cda"></a>

## azure_pfx_certificate.password — azure_pfx_certificate.password / a9cd84034d67 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4)
- azure_pfx_certificate.password

<a id="canonical-6012d8121a30b2c4037ae600f862ffef601d41ba9b1b32adbf87200d4fe4dabd"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad61af9c8f48e291c94f542fdfc9e41f5e7a1ac61ea7a5e017b523f8fe62c0d7"></a>

## Direct properties — azure_pfx_certificate.password / a9cd84034d67 / 3

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-740572c7c12a6ad7984148dd8c929048c8a15c887d56444d0c6b41203fca6475): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-088ca23d1289343ef1b9ef9c7d107ddb2f2bce48a4c890c6d02682e4c28eef82): complete subsection reference.

<a id="canonical-2464c53542a5c59f785115a57891c3eeab692f4cf346fb83dc1d2e46d3eb73dc"></a>

## Next pages — azure_pfx_certificate.password / a9cd84034d67 / 4

- [azure_pfx_certificate.password.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-740572c7c12a6ad7984148dd8c929048c8a15c887d56444d0c6b41203fca6475)
- [azure_pfx_certificate.password.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-088ca23d1289343ef1b9ef9c7d107ddb2f2bce48a4c890c6d02682e4c28eef82)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-740572c7c12a6ad7984148dd8c929048c8a15c887d56444d0c6b41203fca6475"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c96c27a0d1de3bc6623be991b5025eed5731919b8668999f1fdd2e60943c6b31"></a>

## azure_pfx_certificate.password.blindfold_secret_info — azure_pfx_certificate.password.blindfold_secret_info / 5a7ccfc2c0fb / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4)
- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d)
- azure_pfx_certificate.password.blindfold_secret_info

<a id="canonical-dea3b516b0147de14036ccd2739f1a14f0ef9b4cf574909ff38615a613fd7502"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-43da40a61f1263ec15b02257d0486adca3268a4e95c72604a658f4839b5d8b79"></a>

## Direct properties — azure_pfx_certificate.password.blindfold_secret_info / 5a7ccfc2c0fb / 3

<a id="canonical-3008d1320b299b99b3ba36d39ad1dfad2f878dd966fb3dcf425f243c2b68e0aa"></a>

<a id="canonical-62532f6b71536f2bda808eb89e8cae25c999f5a5e3b7904a93668f9ccd95e384"></a>

## decryption_provider property — azure_pfx_certificate.password.blindfold_secret_info / 5a7ccfc2c0fb / 4

Type: `"string"`. Optional.

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

<a id="canonical-144511479c49e621748778433d201504bd37ab0a61d9d1a1ecf76cca13e7d28e"></a>

<a id="canonical-c29b20d8527f145f06a96a09d524a828563455a7162a91e6d0f80b9a9268ea87"></a>

## location property — azure_pfx_certificate.password.blindfold_secret_info / 5a7ccfc2c0fb / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-cfe18e7a64b37e5b8a5bb3c0c33706821d8ae62e48333e611117efc1d2f4b8b9"></a>

<a id="canonical-fb8f34d1f97ae8ac932e4a70f50530753824db29aca59c3452e1a76b9397654d"></a>

## store_provider property — azure_pfx_certificate.password.blindfold_secret_info / 5a7ccfc2c0fb / 6

Type: `"string"`. Optional.

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

<a id="canonical-e44d4873e1a73897285a2d2e37e0f0f5aabfdd2ac01af6ff7bfbecea6f9e3089"></a>

## Next pages — azure_pfx_certificate.password.blindfold_secret_info / 5a7ccfc2c0fb / 7

- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-088ca23d1289343ef1b9ef9c7d107ddb2f2bce48a4c890c6d02682e4c28eef82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dead695ee101484e4d318ad3ff4cc50f6c8c4b81305c0e4b5c92f92a74d36de9"></a>

## azure_pfx_certificate.password.clear_secret_info — azure_pfx_certificate.password.clear_secret_info / f319f6a8ca7c / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-d9670a783b6e4c58fe4d91841ffcd56ced56d310123339ad0292a40f27a296c4)
- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d)
- azure_pfx_certificate.password.clear_secret_info

<a id="canonical-8ea8654b2db3f2bac34109e7b363b31db9ecd1cd0f02c6fac79a34a33c3197be"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b26a0c73bcf6d21bb79206e87182b2473511d85be954f4b91f34658125e53d9"></a>

## Direct properties — azure_pfx_certificate.password.clear_secret_info / f319f6a8ca7c / 3

<a id="canonical-93fde45572ba704d7c341713d8b54842333bd3871bfc5139f54982f916e8e7ef"></a>

<a id="canonical-ae4cff2fbffc307d768473ca7fa787a89866a0a5eaaa6614d30ab75bd60f76fe"></a>

## provider_ref property — azure_pfx_certificate.password.clear_secret_info / f319f6a8ca7c / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7f868bf8f38d544528f4423a4c55b56689dc8ad466a13174e5ed4e3020a19409"></a>

<a id="canonical-f459c7aa10683d25e424479c425e4bb25082f561d9af317fea013ade9a05dc4b"></a>

## url property — azure_pfx_certificate.password.clear_secret_info / f319f6a8ca7c / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-e36fb0acb89af48ec7e2a0406a3efb9fcac970fa5ceb153f1ca94f1aad5e1c04"></a>

## Next pages — azure_pfx_certificate.password.clear_secret_info / f319f6a8ca7c / 6

- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-8d737cfe7746a75b87f43afc24f4248cb9e506cb5af9b01727154cd71069e22d)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1de30643795d178ba1e5bc3ecf7d730910c72e85714f9c11ceacd57cda6c286"></a>

## gcp_cred_file — gcp_cred_file / 0f086282166b / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- gcp_cred_file

<a id="canonical-c6d5e0047cb0140207434ad4cacf0c414ded9691a9e46f5ae652d6d038bd5e86"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
gcp_cred_file {
  # Configure direct properties listed below.
}
```

<a id="canonical-1fef13d9c2499ec47e18169833308e93e28565e1bb08fb5c8f38de9fadabc232"></a>

## Direct properties — gcp_cred_file / 0f086282166b / 3

- [credential_file](resources--cloud_credentials--reference--group-001.md#canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14): complete subsection reference.

<a id="canonical-c0f330ef277a9b5afd260f193952b23ec492b47a328be5299339ef1abb5060c5"></a>

## Next pages — gcp_cred_file / 0f086282166b / 4

- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-659d57915358a89e96680f9be570c2768708d238897bf1ece1e92657d82cd9d5"></a>

## gcp_cred_file.credential_file — gcp_cred_file.credential_file / b9ef5328b737 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7)
- gcp_cred_file.credential_file

<a id="canonical-112ae79be7b74f472b49b8bf1962c588330c1d56f688dcc30a796c8683cda5b7"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
credential_file {
  # Configure direct properties listed below.
}
```

<a id="canonical-8b675f27c698942168df753186a5949a90f9b38bd53bcffbee6f225bf7f101f6"></a>

## Direct properties — gcp_cred_file.credential_file / b9ef5328b737 / 3

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-c1f549e505499741cca89667658e8ba9d8cc0054f0233f426a2d10f1972eec58): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-2b96eed5593a484f1c89b12843795f75a613cf2408fe546a7a87f29b845be5b8): complete subsection reference.

<a id="canonical-0942e6f2cc39eef8265cabf0b6124171e6dac2684a74368f424a79d266d87483"></a>

## Next pages — gcp_cred_file.credential_file / b9ef5328b737 / 4

- [gcp_cred_file.credential_file.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-c1f549e505499741cca89667658e8ba9d8cc0054f0233f426a2d10f1972eec58)
- [gcp_cred_file.credential_file.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-2b96eed5593a484f1c89b12843795f75a613cf2408fe546a7a87f29b845be5b8)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-c1f549e505499741cca89667658e8ba9d8cc0054f0233f426a2d10f1972eec58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2259379daed02f5624e71eb6f2d7b7fe37de0ba21b46260121b7818a7df61a5"></a>

## gcp_cred_file.credential_file.blindfold_secret_info — gcp_cred_file.credential_file.blindfold_secret_info / f50ea21327af / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7)
- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14)
- gcp_cred_file.credential_file.blindfold_secret_info

<a id="canonical-b73371e9932f7968ba9b9f87dce78fcc42b4749cf0f091060650c1207f142984"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f3509fbe355330deeb0345135db952757e8bb3f5acf40f21cf97b60322c75bb0"></a>

## Direct properties — gcp_cred_file.credential_file.blindfold_secret_info / f50ea21327af / 3

<a id="canonical-ded1bde5f500a00b68faf3964668681561389ac608cfd29333883e69de64e9fe"></a>

<a id="canonical-a18756a5c110f555507bc59796922077189ea5cd28e8279e0309be0cf4b21fab"></a>

## decryption_provider property — gcp_cred_file.credential_file.blindfold_secret_info / f50ea21327af / 4

Type: `"string"`. Optional.

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

<a id="canonical-0741dafc23007cfaaff70fa0a316f14f662e597f96d3b17cfb57b67b007f4020"></a>

<a id="canonical-82a5ad64b7fb7e65ead7cd096c0bd09d7dd62307a6e9b79f23ac7f2851b1cae7"></a>

## location property — gcp_cred_file.credential_file.blindfold_secret_info / f50ea21327af / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-dfa60a07e3198c82c322cf47b2ec05614caa1a7c9377bad249bea41c1d4909e4"></a>

<a id="canonical-d8d975f7790b97b7bc29ab4b854026442ad6580327a6ece9705c033800840c5d"></a>

## store_provider property — gcp_cred_file.credential_file.blindfold_secret_info / f50ea21327af / 6

Type: `"string"`. Optional.

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

<a id="canonical-0e864cf8cc14f4e15bf11f13f8e10efccbb62a32d43a7792f65ede3e418c5041"></a>

## Next pages — gcp_cred_file.credential_file.blindfold_secret_info / f50ea21327af / 7

- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-2b96eed5593a484f1c89b12843795f75a613cf2408fe546a7a87f29b845be5b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1cabce2bdb7f6ae41ab074764e3634e2c23f5e25cd3b39cf1ebeb6b709c6e54"></a>

## gcp_cred_file.credential_file.clear_secret_info — gcp_cred_file.credential_file.clear_secret_info / 28694bf8f172 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-eaf28e85327288aae51c56bc0cc7f4eabea47a39fde34da09e985726026564e7)
- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14)
- gcp_cred_file.credential_file.clear_secret_info

<a id="canonical-c9b120cbfd1d032cfe4cf254f8b6ca24eed2f503d1aa499a0903aa3a4a4fc75e"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-793356f00a38d78058bd7055b3e8ae7686645decbb49f6ebd6704680591db4e7"></a>

## Direct properties — gcp_cred_file.credential_file.clear_secret_info / 28694bf8f172 / 3

<a id="canonical-76cb4f180384af9f43aca561452bbc983599de9fdff772d921458d0dfabc0c5d"></a>

<a id="canonical-119292b933bd91228d392977f2b92bdc26a9d5e76cdf6659a30f7fea61a99b0d"></a>

## provider_ref property — gcp_cred_file.credential_file.clear_secret_info / 28694bf8f172 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-139f8e150e8996473fa4c6a9c13af0fc64d971830e3bdec4290470ae96724926"></a>

<a id="canonical-994b76e88b7aa3ac54b481e810ce0dc06113078bde1a3e89d85529c5ff4a909b"></a>

## url property — gcp_cred_file.credential_file.clear_secret_info / 28694bf8f172 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-cb9252340b8190d68a8e24b386207f596bfab1e19290e2d70d6bc7818f004d85"></a>

## Next pages — gcp_cred_file.credential_file.clear_secret_info / 28694bf8f172 / 6

- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-4c7067ff2a6a1df6617939577c372388bb8b4f00c063cf57839aa8a31f040e14)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-3acc6c96cd5416a3e10b6797715d00081c043429c9034ec56cfd26dcf256fc19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-131b057433e7a76f58e9be54caec83219f96dfe976b00d1ffac1b960c0a488b2"></a>

## timeouts — timeouts / 8d7b303e32ba / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- timeouts

<a id="canonical-44c9716c36756ca2eace2ede36b56d0fe85dd0d549e848b75bb5a838c5abf87f"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f804b75a502e7ee8793f38eee1cdc436e61dbc95d87e70419ec6bdf461d4cfbe"></a>

## Direct properties — timeouts / 8d7b303e32ba / 3

<a id="canonical-dbf5f1b102f11a69bce280da5449e40f442c2ad59f27b9b34127e96df1aa73d3"></a>

<a id="canonical-c2245f9825692ce6c199810656f6441989ebb318a5f2587690e8b4ec4b061f88"></a>

## create property — timeouts / 8d7b303e32ba / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8c4eb1fe1f24394877fd2beed9dfc1131d5bbdfd2263553890eead22ca0f82d7"></a>

<a id="canonical-ce3f81abd42f68f17b860aff4a3ddefa09b5405ec3fb296f381e6be1a000c452"></a>

## delete property — timeouts / 8d7b303e32ba / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-d48b89934793851663916b540810463da14d022e42ed89cfd3d6bd185ca05cb1"></a>

<a id="canonical-276096439009d59e0954b259c84dd9cb21d4923301569ff1cfe807cf6a9bc788"></a>

## read property — timeouts / 8d7b303e32ba / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-de2c35d1fbb4fb4407a978591fefcf8c73defbaffa1ccd59191405871e2ad3ae"></a>

<a id="canonical-b9a91b33fb36b44ff3e997627e12c8918a032bcdda3b2e6afa2e123d80707f55"></a>

## update property — timeouts / 8d7b303e32ba / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ce35e80f179a0845358f0f9ba3a64d5d3cb9a6d25f42deb5c5fbda657c03dc79"></a>

## Next pages — timeouts / 8d7b303e32ba / 8

- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
