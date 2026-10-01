---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-c9dfe0f867b8b6e56700e805128f926b5ae3f9e52c2df68e0307ab101f6acf17"></a>

## virtual_server.https.http2_client_profile — virtual_server.https.http2_client_profile / 570ebaffbf6c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.http2_client_profile

<a id="canonical-b0d6cc993ec4cd1c5c60643897fb53962932f49c554aa1170f2b124fa75ccf14"></a>

Type: `"list"`. Computed.

HTTP/2 Profile Client. Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d4655a7389a811e9e1c14edb232af37aa2051c67d3e891ff3a13b93194b5cdb0"></a>

## Direct properties — virtual_server.https.http2_client_profile / 570ebaffbf6c / 3

<a id="canonical-a3336b87d9ba29ea1c83298d7ebc0ceadf9f4dee7822bc1dd3b7f9a59d13bc53"></a>

<a id="canonical-6abc791e197b7a071789bf3a5e5ccfac155cfb7a727b247f2a4693fffada80e6"></a>

## kind property — virtual_server.https.http2_client_profile / 570ebaffbf6c / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-c818848fcd6fbb9351fbf64df48d33c03b5f70d2ad3e842d58e22a2d09403483"></a>

<a id="canonical-aff3f05b04a0befa62a46d3c30ee64eadf5073ff362e1279055aec9c7ab6076e"></a>

## name property — virtual_server.https.http2_client_profile / 570ebaffbf6c / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-8d40102d4e443c6a5c6634be9efcb7447745dce6ff63d0d7642b93c1b46a367e"></a>

<a id="canonical-18bd25bbbf3243035aba4f18ebb7a56a774f2e883b925c9bd451c7afe3917cdc"></a>

## namespace property — virtual_server.https.http2_client_profile / 570ebaffbf6c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-92ed7ecef32a7ce4d4161521d2fce018e0bf62c42a8b54b9f9819992bdd6c20e"></a>

<a id="canonical-c9c0c026c601459d0f2e9da5930a8f7dc772d45c240d8bf1c4c3a8a7368ae044"></a>

## tenant property — virtual_server.https.http2_client_profile / 570ebaffbf6c / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-17ad9c332c2f7c621f7dfdd59723fddcc56aedcc8e2e65d7d08fc9ee200510eb"></a>

<a id="canonical-22836046ce65790aeac9f0b12f838c107e3b435bce66c338f91dbcdc52ae0844"></a>

## uid property — virtual_server.https.http2_client_profile / 570ebaffbf6c / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-bf9dcf104742eb921754c66c26c2b52ee9787aad652f8ecb7b8ea4961f87ab85"></a>

## Next pages — virtual_server.https.http2_client_profile / 570ebaffbf6c / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-1f4d56a60c7aa9907aab9fa958c49f1db3289aa085bc89f77d289cc78512d70c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cc8d5079d2a21db685f01da2c3f6bd038491c6ad285bb73a49a7c04ba262db3"></a>

## virtual_server.https.http2_server_profile — virtual_server.https.http2_server_profile / 2ed2027348c4 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.http2_server_profile

<a id="canonical-cedc1b1efb6532ee0cc2414c5cccb67da37c211833832a853225c077daa91cd1"></a>

Type: `"list"`. Computed.

Configuration parameter for http2 server profile.

Upstream description:

Configuration parameter for http2 server profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-804205282e6c9d9555302858ecfbb52d315db4ed4d496098e30d3723e81f1e60"></a>

## Direct properties — virtual_server.https.http2_server_profile / 2ed2027348c4 / 3

<a id="canonical-62c5e9fd6324cbeedcaffad5ee49cf2dce181a948c92cc5150753214e14c7b19"></a>

<a id="canonical-03550279a933e645ca966b17ad894fa4f76c8231fc4b0199dde100090925c89f"></a>

## kind property — virtual_server.https.http2_server_profile / 2ed2027348c4 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-ebf2dc4990ec58933361adaf5c759f97f2915aa634b6e657012990423742c840"></a>

<a id="canonical-6a1e3c02a97665b3b2ca557d42867e549cb8ca9805ea69d72d8ecc03aea297ae"></a>

## name property — virtual_server.https.http2_server_profile / 2ed2027348c4 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0fbb61df015bd9b18f41cc6bcbcd8b017de0656e2ad748c1531a1b73d20a0412"></a>

<a id="canonical-f7b07d6e1c431f0dd4781e174b35a1dab41cd7023e2f7e63ec85139e635ca438"></a>

## namespace property — virtual_server.https.http2_server_profile / 2ed2027348c4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-d986d1d0a1ee7a4197d3a2902eecf6da6e7fd8629f0e6cfca9e19f45107dc0df"></a>

<a id="canonical-c645c85f5cc646d31f13359a8abcf1ad8d140ea9b6d2bc76c4cb0cb5cb165d38"></a>

## tenant property — virtual_server.https.http2_server_profile / 2ed2027348c4 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-028f4b9d1f7ab90ef983ed22c28e421d233a5926009240cdc5ea151f2db4f805"></a>

<a id="canonical-957d765529fc3327cbaeb578fcace2a233e0713b607dc7838291e8984dcf362d"></a>

## uid property — virtual_server.https.http2_server_profile / 2ed2027348c4 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-af9862aaa86da9ed6a7a5d4a138a57d596df47f434f2d0b135260574c541775f"></a>

## Next pages — virtual_server.https.http2_server_profile / 2ed2027348c4 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-4f2295adb60b5954627ad011694f3a310228b53f7068dabfe59974cabeba4f40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4649fb877fef9d002bb09819c9fa4fa9baf50e3b12a5f51019d47ec518326326"></a>

## virtual_server.https.http_client_profile — virtual_server.https.http_client_profile / 432f0fe52086 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.http_client_profile

<a id="canonical-ddc17c08055c18109993a607e59e76af0393225f3c228b43fdfa08c667e834c3"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ea9da3686dbbae629202aebe12765d95289ad0ceedf1cbf4f9db7fcf26fe39ca"></a>

## Direct properties — virtual_server.https.http_client_profile / 432f0fe52086 / 3

<a id="canonical-35b4daf643bd40497ce7c4a1f405ae122fa0282d491803532b0fe4585fee388d"></a>

<a id="canonical-ad12bde8794a00d4cfdf056cac8a9684282d890d1ed0d6bbe775c25b6e11faf1"></a>

## kind property — virtual_server.https.http_client_profile / 432f0fe52086 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-e4622718da81b6ab913640e16d2177535859cf25b7136cbc2b649e5a583f9c46"></a>

<a id="canonical-f99587b15f42f380c0a09a794c45bde0a589288e70ed63debb15d336329fdd37"></a>

## name property — virtual_server.https.http_client_profile / 432f0fe52086 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-17e385808a2277afe1ee631ee68a115849edf62d5cf1b365dbb0039d5f149ea5"></a>

<a id="canonical-80bc7675058722d7081ea0bff11b62feecb9a7291e70755566a22bc92397d4e3"></a>

## namespace property — virtual_server.https.http_client_profile / 432f0fe52086 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-c769e3a8f0816dd99fc5da9bd0bd91f7bdb50ec6d1500048cbfd79e6913101c2"></a>

<a id="canonical-fa17e8a471a3467df4aee2e2369f1ae49b4ae0f81b7dac0af966652130afc6a5"></a>

## tenant property — virtual_server.https.http_client_profile / 432f0fe52086 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-26b0d1062156ab0972b57d5edf6dd2dc09641b5b5ec66bce962a18960029ec6c"></a>

<a id="canonical-64c417af336c94e526af111554f4c7666798e3a80d7d5f9314218cc1004c9535"></a>

## uid property — virtual_server.https.http_client_profile / 432f0fe52086 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-d9fc2ab6f3635e1a1de80736e3e45e72ae1afd0e4596169494db5557d4a4b2b1"></a>

## Next pages — virtual_server.https.http_client_profile / 432f0fe52086 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-5ab92633e91ab2b2e5e181271c742ebb366df4afb6db1748884f84225fa1f911"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff82a248c0f0fed0429d7e374f532d90ab34cf54a554a2cdafb51dca38b5b480"></a>

## virtual_server.https.http_server_profile — virtual_server.https.http_server_profile / 6503fe220644 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.http_server_profile

<a id="canonical-dec9eda488ec8a960180066b26bbca0c038abf43a9f4f014cfcd89592d076588"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cc893217fb1f7b718ce5aa61e4031b4497e4adb4037048052a76d77c438ba744"></a>

## Direct properties — virtual_server.https.http_server_profile / 6503fe220644 / 3

<a id="canonical-b76e29ccfd6be22a2d7752b58ca78ce4d53ffa95b14044807368888061bb6eb5"></a>

<a id="canonical-fc0b24810e6889c2954b115adb3c9199f9cb7ccdc36af12d2213f283371d52ee"></a>

## kind property — virtual_server.https.http_server_profile / 6503fe220644 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-4f7de1a837c82b7e3afaef25c7713766dff443c197e97d32f914048d830d1f67"></a>

<a id="canonical-28a4abc100f6ff7461217da9e79c74b4d12ddb2028cea2c28730d51f7ec92ba6"></a>

## name property — virtual_server.https.http_server_profile / 6503fe220644 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2364178b23216538d2b01b45afca31964db33584e28d302b271f6ed2119a8fda"></a>

<a id="canonical-b87b46070ba85d0e7f26a11c8f369ac8bfcf013a0d68f7a110964ab7cb413180"></a>

## namespace property — virtual_server.https.http_server_profile / 6503fe220644 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-a4acbdaaee593c4bc03e5814eb77e408c1cd453d0348661712fe55341692f3ab"></a>

<a id="canonical-53eefa8729f645c50156735c3bab84eea991aa4531047771e17cf1a20b66eb75"></a>

## tenant property — virtual_server.https.http_server_profile / 6503fe220644 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-46bb3adc75666cc825e48237c0b0d3c51b2d541412251fc10571f28c66d1418c"></a>

<a id="canonical-baae535ab1f9302cce2dbbb1c42d7d9529a4767f40a50147aa28bf065a8755cb"></a>

## uid property — virtual_server.https.http_server_profile / 6503fe220644 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-904c2783e0211819a14df291abe0fc1d13299dcef99e6d2b51340b41fc262f2b"></a>

## Next pages — virtual_server.https.http_server_profile / 6503fe220644 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-fa2f89657d2a7c991f8eaf934c50ff599ed05d06e1368f775301b07573ad1824"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e11f87d76313d68411045977bdc480449d34574bfa915e6f851a966147f7af8"></a>

## virtual_server.https.ocsp_profile — virtual_server.https.ocsp_profile / 6813cb41237e / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.ocsp_profile

<a id="canonical-87091287bd839c53798f2c02171540c3af9a6338920296d53c343792561bb449"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a4173ccf50a091afcee41a427d940ee5029dd4a0b9f94e48dc7f917797356b45"></a>

## Direct properties — virtual_server.https.ocsp_profile / 6813cb41237e / 3

<a id="canonical-3f78313b58f41f5bf4ea7edc013bd9b05f61dcbbcf677d5825aea5acec2adfe2"></a>

<a id="canonical-6a93499edc9aea8744b8da8248665ea671922e9de3e2a101cff2693279b69cca"></a>

## kind property — virtual_server.https.ocsp_profile / 6813cb41237e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-ac2f76107cd2fb8478b7cfaf13031b8d1c2d7172ae79687846a381c14eb8d6a6"></a>

<a id="canonical-2b78528ae52f8aaa5a8bcb948209b784cd2090ea127419e219a4cd1461d204a5"></a>

## name property — virtual_server.https.ocsp_profile / 6813cb41237e / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3d68d3a18bcc911bbfd7e8178558d6cef2ddbfd706bde3df4c1b3329e85d883c"></a>

<a id="canonical-1c2770448a2e9b14950b6003ed5c1ac60a37e537c7ef5937754976c3d027041e"></a>

## namespace property — virtual_server.https.ocsp_profile / 6813cb41237e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-fabf427876a2f6574836ea4be14cd382e43fc69447e6013b0043a7a2071bfa52"></a>

<a id="canonical-c67d1e8d7e599085cb32f4572118cb037297dc0a660b41bbde3b9557893d6cdf"></a>

## tenant property — virtual_server.https.ocsp_profile / 6813cb41237e / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-d4cec6258e0adaf178491b2082030454545e59440cc2a24a6a59f1e07001c360"></a>

<a id="canonical-d907f9cf46de1536f0d04a5f2d83b6d8a9e2d14bc601d8d86a8957261c72921c"></a>

## uid property — virtual_server.https.ocsp_profile / 6813cb41237e / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-b6ce65e15e9c2a33f0cbf5a7aca74d3a36aa6921fe9726c0af8bdaaa6f877c0b"></a>

## Next pages — virtual_server.https.ocsp_profile / 6813cb41237e / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-6e7660e2ef830fb2c83dd4360748cdcbef6b9b4a757cc62fd62105258fdb1cd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f822399211d958d30a05dda13debef3a1e961d7ddc477d0329e0b28b813fea00"></a>

## virtual_server.https.server_ssl_profile — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.server_ssl_profile

<a id="canonical-f3d6bf36cffb58d7a7dbcfadb032655eae87008585c63f86edff665e168906a2"></a>

Type: `"list"`. Computed.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7c962bf5db15d6375845dd14250efeaaac3c9c09ae706f31be24f35ea68d6c49"></a>

## Direct properties — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 3

<a id="canonical-83603a426b70080275a6d33cfb1fba108165ea1d12ddbdbe41a3f7c1c97df634"></a>

<a id="canonical-b5416b27057aecaebc1af552a3a52afea02181f8f1c15adb41462df54b4fbf63"></a>

## kind property — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-7bb36a8b523c3bf3cf5e1e83b1da3fa468b39d3a8dfc6fc539bfe6d7d1a33cef"></a>

<a id="canonical-92a1ef7a4ed173d52e4b445c28038bc5b7226c9b30ca7c87d65079509d23d8e6"></a>

## name property — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-00628772f6e9a747acf4e45a81a99bb8eb868a55ed6536a1fe29da10b0afb449"></a>

<a id="canonical-2ce7a8280dd58f5bf78e93da184d2413e6ac8759ba6909cc58eb757260773098"></a>

## namespace property — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-e3f1ebaf5bed7745a5919e2612cd8546e09858fcdafc36dc6b9f92d8ea47395f"></a>

<a id="canonical-914d1eae383fd38b62894539241fa6a5c01a1776ba247f4a819ce3c08b71a4b8"></a>

## tenant property — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-5dd28854200a7a322569ae81d9550a844910261ca01effcf856977b0b0219e2e"></a>

<a id="canonical-2e766236068c45f861255572c8c50bf44595d827fdd49b3bd7d4176ecb1e147d"></a>

## uid property — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-a0ab1e10f926903deaa4b849361a20c6205f5829b17223f788a33fa0cb36f258"></a>

## Next pages — virtual_server.https.server_ssl_profile / 44acbb4b4063 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-240a2b23fed325227dd634c21c987962bb8f31defa13b73f33971e6832a5aa13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ca5f3cbca6e9a097b6379cb18019dec1ca16231cfb1cc86ad8c715e26ff775b"></a>

## virtual_server.https.stream_profile — virtual_server.https.stream_profile / 4ba19b25e8da / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.stream_profile

<a id="canonical-9159b8d42da7929cd3846360b09881856c3ba471f5eb18c05241202f3da79dc5"></a>

Type: `"list"`. Computed.

Configuration parameter for stream profile.

Upstream description:

Configuration parameter for stream profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7e879c9370520dfe0e1710128eff9c9cf7ed4debf0c4e324083a9c6137a96d31"></a>

## Direct properties — virtual_server.https.stream_profile / 4ba19b25e8da / 3

<a id="canonical-99a9ab68c36646f9ca9d0cc2964aa93d1a52e8e188504b07f48b9c4150967fb2"></a>

<a id="canonical-69fd892c6e91ce7ce5ed7f297ad02700201229292fc2405404e0ae55a3ee1167"></a>

## kind property — virtual_server.https.stream_profile / 4ba19b25e8da / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-6994a2a8b89afcd166c194b0791bd49b031690f34f2cd54edb690662a089713d"></a>

<a id="canonical-851fe1691232049735dcfdcec6fe7f0b96f2ca100c680884a9c8ac4df416eef7"></a>

## name property — virtual_server.https.stream_profile / 4ba19b25e8da / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-7bc71d00a060e9fc6f82e242d247f5430e5f4493fa1f58437f932043e43194ad"></a>

<a id="canonical-e9d3f09959891a174c6a275feb69318533cc59e786507ea5b9c1b6d33cdf3d11"></a>

## namespace property — virtual_server.https.stream_profile / 4ba19b25e8da / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-d616a5c45083d03e307ff5bac81328340a7e78178a2eeeabcbaed5e12523614a"></a>

<a id="canonical-c3dc9c943d5d66590b5e419eda9420375f3525d3739044d9ec38e0c5aa2b4f90"></a>

## tenant property — virtual_server.https.stream_profile / 4ba19b25e8da / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0e6fffec07b17a047715a9fea8bd75f8bbfc0733e26040c5a532bd4dd2c41306"></a>

<a id="canonical-088ac8bcc223c4b75b1a9911e8b05d70f0168e8972adca15e319ea19705381ad"></a>

## uid property — virtual_server.https.stream_profile / 4ba19b25e8da / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-c6fc7336fe86f8d6f6d616d85ab375efbb6e4bd3e5d6d90a19bb2ad54d8826e2"></a>

## Next pages — virtual_server.https.stream_profile / 4ba19b25e8da / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-a2de95bc4627e2378ebccf89c853098db318d0565882d5ccc198ffefed9bbaf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fbdee55be14760950112f26fa35da5ece9296fd13701735f52fbfb322988b09"></a>

## virtual_server.https.tcp_client_profile — virtual_server.https.tcp_client_profile / f601de22da79 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.tcp_client_profile

<a id="canonical-f542921f0bb847a58417d6670d00afc58412f338750f7d3aa4e6b61e839cf44a"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5b866bfbba394971a8c8f5222c62ffc5d5e2f74728d66e6431300617803a791c"></a>

## Direct properties — virtual_server.https.tcp_client_profile / f601de22da79 / 3

<a id="canonical-c01d32c4bc9b4ecffd848aa0b03821594c54ee4cdbbf8f38580e6dfeebf3faed"></a>

<a id="canonical-1b84fd60d9b0e552c92a8d50e914d4a46221b06bbef733c97a9975f94051a49e"></a>

## kind property — virtual_server.https.tcp_client_profile / f601de22da79 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-5316816f398975096f0329db7c3983714e4d64a0e7d526349d06897976b61e99"></a>

<a id="canonical-894ea8915661b92bbcbde02a865dd3b0e6eb0316ea9f371af506dfc6ade776d0"></a>

## name property — virtual_server.https.tcp_client_profile / f601de22da79 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-ee10a6746546badae5342c16a921a16babbb7288400cbe2ffe9570565f050d2e"></a>

<a id="canonical-416d76814dc39e2161972f8bdd383b71241caf097afdfd8e20b2b0097ae5bdbc"></a>

## namespace property — virtual_server.https.tcp_client_profile / f601de22da79 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4489a866eb442ef5fb1c5cfea10f743c02f678e2247e5e9d879769bf04b477f0"></a>

<a id="canonical-f0a8277b8b99e431bd08a42c2dbe0b63739d18d332641cbe8eb310dffcd2fd65"></a>

## tenant property — virtual_server.https.tcp_client_profile / f601de22da79 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-08e81b4f088414c05c1ed6c9679c5502b7c10049e2ac533243edf52f90197d4e"></a>

<a id="canonical-11856a2e723281b22baf33fdfe4dd5a646970af8500f44feea8b2cc6daadfe3e"></a>

## uid property — virtual_server.https.tcp_client_profile / f601de22da79 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-f67ff9e1cb2b5d574667f40c3b021e06bdc34921b536be98bffff04ebd963b65"></a>

## Next pages — virtual_server.https.tcp_client_profile / f601de22da79 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-bb93372f0965c7716127d587edd71646800d126364cc343c2c10f4c0a4476d63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a888b0d60c9942152219e25676aa4aa1c9f3f6627f246774c61aafb7120d94dc"></a>

## virtual_server.https.tcp_server_profile — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.tcp_server_profile

<a id="canonical-bf6a3b12fd0251570d2fee396000d411ff7e7e4a1d3c82a7b1af949c75317422"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-193db92b9c10738dd913132191d71eb03fa3b3c2ecc5631ecff95d89779831f2"></a>

## Direct properties — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 3

<a id="canonical-071bb0b692f5bb69f9a93462eb9915f201a3963077ce3b28d09b99f28e706fab"></a>

<a id="canonical-325ef6937fb485f7d1958b4517453f1e227fd4e69740166ae9c2449648a243d4"></a>

## kind property — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1111f87675cd6a6437454f231949b2e0351c2560e0b6916efcd63fd040e21d91"></a>

<a id="canonical-5206b0cc6a9f1c06f0adbdcd2e9c70835ce71b1eed972071d2d2b5fce0f5d10a"></a>

## name property — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-e753c24a70ce44f56071ba03abcd57d00a33bc69f38e1b3350a53c0dde321d22"></a>

<a id="canonical-5528858b4fcfbf0ef14b13201816c281e65bc85d1b73c794b4fdd19bda580a9c"></a>

## namespace property — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-fee442f391679ed3caa0e8acebb1c571b5d1d381db32ba0fc48290a22387c60a"></a>

<a id="canonical-b6757a201d1485425c3b3fceaf939ff40ae6f95e53db7e6738573d95a1f3af78"></a>

## tenant property — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-6bdc4df7bfa9cd2dcb14a08ae310f0a92158700752c5a23cd535b76c20f32cf9"></a>

<a id="canonical-855523c63ef8892e1beb024b77ff859e8d66a62bcc6dd95b0a17df6dc4ddf146"></a>

## uid property — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-a209dab841a9671eb6f8bf5a09acd55fd1e649a93b98e82c24ea71deecd30a03"></a>

## Next pages — virtual_server.https.tcp_server_profile / c66f8eec20f1 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-93eedc7c0076712eee02e8427f787d05503b61bd07efcd3c772c90a582b01c04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46bcd6782ed4ee215ef652c80b0d82da4679f7c2773d8c31ee82f37d078a103b"></a>

## virtual_server.https.websocket_client_profile — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.websocket_client_profile

<a id="canonical-d458aaa7cd4c3b57a43d0f2b347b04de2bf6ea2c0b8abc31a7bbff84b0e7d77d"></a>

Type: `"list"`. Computed.

WebSocket Profile Client. Web-related configuration

Upstream description:

Web-related configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-563d8ae657e1f455a6a386790bdad7e8e4025558f9fa926ae47fe0343a60576f"></a>

## Direct properties — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 3

<a id="canonical-9d3bd9350990ebcda6514edc832838f02897ff096719f1a6b7eea75260620ee2"></a>

<a id="canonical-93360236369f69beec33c65df7c4e1644a90c00316dcee4a3f22c872ad125696"></a>

## kind property — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-fceece888635c30c9763cc902782df97470d8478af7a82e7a70f5b1cdf3544e4"></a>

<a id="canonical-e044f20f806ae3317d56589bb61879f5138fc26b40a921e3f203f0e9b149d90d"></a>

## name property — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-148d99bf8efc51daad5728650b6b36459876f3b2d0373266a47a89fc6de81292"></a>

<a id="canonical-66504f871d7645ead6a01669fea5b0f0a3313d444867b29c3267b5a0270d6ea3"></a>

## namespace property — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-444dcf30ef004905b057fb6d028cecd43b6ffd4be4dd2c766a5dde7b43220ad8"></a>

<a id="canonical-1a440088d8b2f84c45632757043890b789a28afdf8c5eef93637dbc3d822d67e"></a>

## tenant property — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-73b93473c0acfe0b4edfeebf82deed9ecb421614162c869a4ee722fefedc65b1"></a>

<a id="canonical-a8d32f7a5b32e3082f6b5cbe121c092bd11157fe22a14038d4303b8cf203a057"></a>

## uid property — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-6ebb2c5459c0e6986773bce70be21daf929296da7b2a7d1c9771c6e37b08aa0c"></a>

## Next pages — virtual_server.https.websocket_client_profile / ea3ddfa399f9 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-c6f803d16e3092136fe7867446111a7a15a6010c339dfa4ef2439b4ccc4779bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8805e97fc464dfaa2af8292b1cb434dd369094ba4d5a510e24b05782445f590"></a>

## virtual_server.https.websocket_server_profile — virtual_server.https.websocket_server_profile / e143b921386d / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.websocket_server_profile

<a id="canonical-86571b7a55e8c33a958a06494cbd77f1519b55ca11d838c124a00332e4f8400f"></a>

Type: `"list"`. Computed.

WebSocket Profile Server. Web-related configuration

Upstream description:

Web-related configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ba4c297f1415e4ccff05a3e620beac62de9fcbd7f677f59731d0269439a060fd"></a>

## Direct properties — virtual_server.https.websocket_server_profile / e143b921386d / 3

<a id="canonical-173ee6154284ade44cd5c10b3c2e43f939e2221017d6eb39054b8ba4ed9c4bae"></a>

<a id="canonical-87ed5e2ac42bed90e7978dafa291d4694dd91d1fcd12ccf76c61b0f9a98e73e7"></a>

## kind property — virtual_server.https.websocket_server_profile / e143b921386d / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-4900917f6ff6e6bc285c2c7742f834071aa0ea224ef8d2427e1ff7c96ed90c4b"></a>

<a id="canonical-3f6c772bcfcfde78b91504e8fb676572077d6ef30714e970c8be2c9876caa0fa"></a>

## name property — virtual_server.https.websocket_server_profile / e143b921386d / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-6da0ae28f2ae011a20cc029d56aade976357b922119b464554786927b396c6cb"></a>

<a id="canonical-4af3757d8d4320e20a33b0c3512fb5f9733b0141a3832f3d3fe7898bde554709"></a>

## namespace property — virtual_server.https.websocket_server_profile / e143b921386d / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-063da4d3c6bc63e7018adfc431f7941f82618485a66b0d3a9b4ad7424d7fd94e"></a>

<a id="canonical-176cc53bf3f39860909dac2548496093e915755d7505eb4e4d8b4b630f561c5d"></a>

## tenant property — virtual_server.https.websocket_server_profile / e143b921386d / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-20cc66aaa78073189612dcbecb93c669666caa87a095416d3d9ce5fecd7b1681"></a>

<a id="canonical-852b13d6fed742ff59a8bcc0f1013a874d8963ca951de3bb08a5279ac8bcdc67"></a>

## uid property — virtual_server.https.websocket_server_profile / e143b921386d / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-1fdda8785a3535b0fe5ab9d9262f7d9604b79d86de9fc9310d9aab5fd4221fa6"></a>

## Next pages — virtual_server.https.websocket_server_profile / e143b921386d / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15893235fcd31d55b8cc132d2b33ebea3a51a0d6c4081f5d84deba66a47c4ef9"></a>

## virtual_server.immediate_action_on_service_down — virtual_server.immediate_action_on_service_down / 3d0be501e198 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.immediate_action_on_service_down

<a id="canonical-24f59946a3d58f9754f08a2e67462e6b9f8a0985f17241d8c05e8bd5bdfa33d4"></a>

Type: `"single"`. Computed.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.

Upstream description:

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

<a id="canonical-d5917c24d50bb4328f5de745e6ffd27607714dee81d82b254dd6e21477a21a35"></a>

## Direct properties — virtual_server.immediate_action_on_service_down / 3d0be501e198 / 3

- [immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-3c52d95f47c16cae0e679bd8142d480b85175debf33b59c03b35b1d234626898): complete subsection reference.

- [immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-38aa031f50155fa2cb249688a803c17ec1565c41c79cc8ec917db163bb509046): complete subsection reference.

- [immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-2b24ece852b837347217b688ed2d5c56e42c18b738949a0af81d7c74bb1fc9df): complete subsection reference.

<a id="canonical-256cdb25f6b4cf48d259bd76b4cdb1b86266ff84075952057978e158ce3f16ea"></a>

## Next pages — virtual_server.immediate_action_on_service_down / 3d0be501e198 / 4

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-3c52d95f47c16cae0e679bd8142d480b85175debf33b59c03b35b1d234626898)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-38aa031f50155fa2cb249688a803c17ec1565c41c79cc8ec917db163bb509046)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-2b24ece852b837347217b688ed2d5c56e42c18b738949a0af81d7c74bb1fc9df)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3c52d95f47c16cae0e679bd8142d480b85175debf33b59c03b35b1d234626898"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb5002d69bf3f5020a0e2dff34f8f90f7fc947c45c87f3c2e0295bbcc496846e"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / e3b4cbbe97a4 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop

<a id="canonical-ef0ed12167543ebba930c555e39a68faba075ef6200abd2eb49872ad6eb8ef1b"></a>

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

<a id="canonical-f4dd05bb930184b30b41b8bc962aa04ce580d36c6c23910acfa682536e2631ce"></a>

## Direct properties — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / e3b4cbbe97a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f5cf27e0f33274e70263b5bc7a7dde509fb43d5487597faf78e2a01cfa4996aa"></a>

## Next pages — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / e3b4cbbe97a4 / 4

- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-38aa031f50155fa2cb249688a803c17ec1565c41c79cc8ec917db163bb509046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5797e257ece716be4c534d16b3abc67bad7c399cb9450ec9c35138eac40050b9"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 740e6a557ab6 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none

<a id="canonical-c0f77183775d8a5679e3fa2bebd2c4effa7f3cbf1ff8c2d2588ea70cff5991dc"></a>

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

<a id="canonical-34b1617ce1fd5dd85464ce16bcdee8e3ca1a4cb8bb280ae2a9dc56b08c7fa68b"></a>

## Direct properties — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 740e6a557ab6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34fde90c63b42cc24c8876f68328c7537e2851e538338b54e1934d1e2213cfcb"></a>

## Next pages — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 740e6a557ab6 / 4

- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-2b24ece852b837347217b688ed2d5c56e42c18b738949a0af81d7c74bb1fc9df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db43f0aa3147de44ae2c18da270ff751be5831ff16b6ad8cea6eb1ea943bf340"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 837affc15ada / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="canonical-cddb9284ff3df29dd074fb333c13bf15d54071e1071e619c775fd379c2dee06f"></a>

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

<a id="canonical-af259bd8b4aed4b395545b227886b4ae6da8bf010a951722622e13fd7e23fadc"></a>

## Direct properties — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 837affc15ada / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22023ce5af00cbc7b72457832fcb7c3cd2ec50ae27603abf713c994a2fa30724"></a>

## Next pages — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 837affc15ada / 4

- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-00c49ee343c420b534167e37120e0f65d3241be0e65da401f8ab8495f7ac87cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c2f886c42ca726430e445c72e7533b1b63448fab8ed89e75ee14941692d02df"></a>

## virtual_server.last_hop_pool — virtual_server.last_hop_pool / 6fb6488b52b2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.last_hop_pool

<a id="canonical-7dfbdf1f7dc13e226b471d7bf9dbc910b5d21e4e2bc8d7b2d406eac51f0b43a1"></a>

Type: `"list"`. Computed.

Directs reply traffic to the last hop router using the specified pool.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6309330c9031974bc355ef7a1475d31007d2c5387ae5f1e17697a27ec45f0faa"></a>

## Direct properties — virtual_server.last_hop_pool / 6fb6488b52b2 / 3

<a id="canonical-f8b09b3656cad01fdeae4a1f7f9fba938bfa425513fd4e2299c6387b7495fafe"></a>

<a id="canonical-a6a5167981c1d6b406dc2e27fee4d38cea8381ca567b6b6550b236afccce7c27"></a>

## kind property — virtual_server.last_hop_pool / 6fb6488b52b2 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-5079228e397112c36a01c694a798a009dae5b3a10f91b0373a1626b22343d26c"></a>

<a id="canonical-5ba13b4de83890f63c38a120ef37ce6e8aa58816699307bf594c780dea0c867b"></a>

## name property — virtual_server.last_hop_pool / 6fb6488b52b2 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-4aafd1a5c6c0411a5f1b251795ecdb7dffa42e155d119e0bab0f74e9e7a3555e"></a>

<a id="canonical-ec2576c42323ea16f347f8d06a94d8148784d08a40bb8a5ba064079dca590972"></a>

## namespace property — virtual_server.last_hop_pool / 6fb6488b52b2 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3a3945ae58fc1137aa54dbebf09b1f882e2df9ff0483c8b30a48c252f6264e60"></a>

<a id="canonical-89ed4829f2b5be1fe14a3e1adba31dbbfafa16c03128066d2dadca566dde1eca"></a>

## tenant property — virtual_server.last_hop_pool / 6fb6488b52b2 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2ab5eb34d4b33e195fb97687719c4088f41a22426fb3bd7da2b656df8c1bbba5"></a>

<a id="canonical-79b91492e731561081f7a7331a36fb9ff83f2d6a95a9c94042b67cb8dcb46912"></a>

## uid property — virtual_server.last_hop_pool / 6fb6488b52b2 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-50963c6677bd7572e2de627412c0df8bf782cebeaf6e212a30e3c1e8da79eaf3"></a>

## Next pages — virtual_server.last_hop_pool / 6fb6488b52b2 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76ba518f41eb66f811f1a9810f0849e5e2fe955a6d2d5057c07fd0a5dba41bac"></a>

## virtual_server.nat64 — virtual_server.nat64 / 6fe1268a75ac / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.nat64

<a id="canonical-d6aa6fe895906e316c1d2f52574869c46bd0aab9012c237b2d86bbcc398455df"></a>

Type: `"single"`. Computed.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

<a id="canonical-5687f451fb4d089451701f1417f31c39d71339b9e8f7cd773803804f7d700dc7"></a>

## Direct properties — virtual_server.nat64 / 6fe1268a75ac / 3

- [nat64_disable](data-sources--application_profiles--reference--group-003.md#canonical-4af7409129696d958eb50dc74a348922f6864e4765648c8be15017fcc37f37a0): complete subsection reference.

- [nat64_enable](data-sources--application_profiles--reference--group-003.md#canonical-874d6a06c6025f04049caf951c2583cda920d5cfbf7e06d1236ac7099c49f843): complete subsection reference.

<a id="canonical-97696c0b3df79c65e66aa3e9ba15dbbb31bd4330347c5d19525691bc3a7edcdc"></a>

## Next pages — virtual_server.nat64 / 6fe1268a75ac / 4

- [virtual_server.nat64.nat64_disable](data-sources--application_profiles--reference--group-003.md#canonical-4af7409129696d958eb50dc74a348922f6864e4765648c8be15017fcc37f37a0)
- [virtual_server.nat64.nat64_enable](data-sources--application_profiles--reference--group-003.md#canonical-874d6a06c6025f04049caf951c2583cda920d5cfbf7e06d1236ac7099c49f843)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-4af7409129696d958eb50dc74a348922f6864e4765648c8be15017fcc37f37a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d982468b04cc5ea50c1a668ea0cca48bfbc8d22bca2fc3b904bc861fd09b6ef9"></a>

## virtual_server.nat64.nat64_disable — virtual_server.nat64.nat64_disable / c7b8c2d15664 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a)
- virtual_server.nat64.nat64_disable

<a id="canonical-fc240bbed6629c5e62589d6f424a8aeb7b4d6eb9c8f0d23144056c7fe4696e73"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nat64 disable.

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

<a id="canonical-960b5b21618e0df4ad68e35906df34186135337e3147ddb9a8068ba20f2f884d"></a>

## Direct properties — virtual_server.nat64.nat64_disable / c7b8c2d15664 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a3795d4956016f5f134d5730c0360c963592982db3dab4bf9bca841d8024afb"></a>

## Next pages — virtual_server.nat64.nat64_disable / c7b8c2d15664 / 4

- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-874d6a06c6025f04049caf951c2583cda920d5cfbf7e06d1236ac7099c49f843"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd2bc221cb2ba0768d8d51a67025b2641b0e8fec973848c7fc6068d66f7b7901"></a>

## virtual_server.nat64.nat64_enable — virtual_server.nat64.nat64_enable / c02debbb556e / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a)
- virtual_server.nat64.nat64_enable

<a id="canonical-bbf30fd71593225aea4a4a36d80ed3f96b08685aa69ae17c0a2e19b7a70fd17b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nat64 enable.

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

<a id="canonical-072a9399f9b586fb4fe8a6e8268cb0f9a08b7759e5db5074e0c2a89d518ce93d"></a>

## Direct properties — virtual_server.nat64.nat64_enable / c02debbb556e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82936b1e34fce715185bb9d4e27296c27ff3906150d21972d0817747b2e59292"></a>

## Next pages — virtual_server.nat64.nat64_enable / c02debbb556e / 4

- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-569da1b18cdfdd3462591b3de77e164dbb75b78ac8227bf04bd0388d1e2c3128"></a>

## virtual_server.port_translation — virtual_server.port_translation / c0d699db4bbc / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.port_translation

<a id="canonical-2c1a9a7dc9518624b1ed03a4160943fcc843e743f8e29e9ef81b323d02be6d47"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance..

Upstream description:

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

<a id="canonical-8b7a823bbb73a9b5bb1c6f23080ba47fa97e6f6e48ba7086f4604b1be4a3999b"></a>

## Direct properties — virtual_server.port_translation / c0d699db4bbc / 3

- [port_translation_disable](data-sources--application_profiles--reference--group-003.md#canonical-31098fb66378d6b2eeb36bbbf283c32e465d73d5eaa901a14b34bc7036658fab): complete subsection reference.

- [port_translation_enable](data-sources--application_profiles--reference--group-003.md#canonical-9ff459843f5a1b01cfb36e5f506a39cd9ea6a0be18de3afaa9a03d25ed6a3311): complete subsection reference.

<a id="canonical-127ee9a2c0f5bb2a3d1388a671fac7b239a10b44783ad956ccc98a265f31771e"></a>

## Next pages — virtual_server.port_translation / c0d699db4bbc / 4

- [virtual_server.port_translation.port_translation_disable](data-sources--application_profiles--reference--group-003.md#canonical-31098fb66378d6b2eeb36bbbf283c32e465d73d5eaa901a14b34bc7036658fab)
- [virtual_server.port_translation.port_translation_enable](data-sources--application_profiles--reference--group-003.md#canonical-9ff459843f5a1b01cfb36e5f506a39cd9ea6a0be18de3afaa9a03d25ed6a3311)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-31098fb66378d6b2eeb36bbbf283c32e465d73d5eaa901a14b34bc7036658fab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06b1ed5cf3088ceed6acba32318a90dfa98cd9c3c7fb91aea87f6f1ac8822725"></a>

## virtual_server.port_translation.port_translation_disable — virtual_server.port_translation.port_translation_disable / 63e634863c0e / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68)
- virtual_server.port_translation.port_translation_disable

<a id="canonical-6bfc5b2bc3deb4a5e7f9b5c5855f9d7f5d87006eea110759fb60ba8e64d5aa85"></a>

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

<a id="canonical-25c87ac9328e33a8e3f96040fb262f8ca7edc174250a3d35eaaa8168fba2fa6d"></a>

## Direct properties — virtual_server.port_translation.port_translation_disable / 63e634863c0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a697c9f15bc2f99a2bca5b1681d23fede55747f8e86e59a7bf7949fea7b43dca"></a>

## Next pages — virtual_server.port_translation.port_translation_disable / 63e634863c0e / 4

- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-9ff459843f5a1b01cfb36e5f506a39cd9ea6a0be18de3afaa9a03d25ed6a3311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c008e594daabd32eafbdd23705cb0ff133777bea426189af63626e85447aecac"></a>

## virtual_server.port_translation.port_translation_enable — virtual_server.port_translation.port_translation_enable / 62932590eaf1 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68)
- virtual_server.port_translation.port_translation_enable

<a id="canonical-5b248fabc30abbbd809c4ca12eb25f1375999bf88acb488f7f2b10461ad1d606"></a>

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

<a id="canonical-36ead42b8152c65861d3b2176ed8bd24f587976c371593104f98be3817648eb1"></a>

## Direct properties — virtual_server.port_translation.port_translation_enable / 62932590eaf1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf8cf56a858bdb3988a4c22caa3861eb2b8c1214e6389e76804f8d06d32f9e89"></a>

## Next pages — virtual_server.port_translation.port_translation_enable / 62932590eaf1 / 4

- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-eee319a7a1bc5aea591a514b66ed5b4f70fbc9ff7220feea6b49aaec53d79ede"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-734cc0a9c3350a079c3c6f2746847224ff42dbc50a7fd9f0ae1adb6ee23697ca"></a>

## virtual_server.request_logging_profile — virtual_server.request_logging_profile / 0855aadc8a06 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.request_logging_profile

<a id="canonical-4d530c7ac5195f5df8e87d097dfb37073fb53f5ac6fc8c695f9aa1f4a8b9c182"></a>

Type: `"list"`. Computed.

Configuration parameter for request logging profile.

Upstream description:

Configuration parameter for request logging profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-76940012a9ee80dcd7da53455ffee6c7449778ff039e4bf09e5916bf90a6a080"></a>

## Direct properties — virtual_server.request_logging_profile / 0855aadc8a06 / 3

<a id="canonical-d516f11b7ab24db2d4a09d761b126baa6926b235afe21cd6e7a41edfaaf4163d"></a>

<a id="canonical-d699951686ddac30570a8a8835afef14554d955e97dee5096a1f5b5943fc5bcc"></a>

## kind property — virtual_server.request_logging_profile / 0855aadc8a06 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0e7499e10e90607a4a3e18de29aefa396cd69047f57b389c42afa59425e4c826"></a>

<a id="canonical-58f91b31ad7b18190ce016c4e847dae974ad652a6f8e9ac11b4757a991a1504c"></a>

## name property — virtual_server.request_logging_profile / 0855aadc8a06 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-bfb6d8aac0620447bf6b6cb044c01e0d7603ff2ec87342f6b03177ea7e2656a4"></a>

<a id="canonical-408257889f68ae2c1039d3acf01342dd28939b24d98291b6d732f36c6256f79d"></a>

## namespace property — virtual_server.request_logging_profile / 0855aadc8a06 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-b9e41430aa0a1d1e36406bca5475c489624d6a2b746154d9768f92633bfdd1f3"></a>

<a id="canonical-ae9529f1dac646c8d58e3e624a6d354017ee8ba6b038e394f2a2087f1546646a"></a>

## tenant property — virtual_server.request_logging_profile / 0855aadc8a06 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-6527c9961712f454e771a9a8daa43ccfe80c9b882ca03aa018b65c61d05f5be0"></a>

<a id="canonical-687ac975436862d4ab55086c8db4336e21b425df8762cdd9259015769a3205fc"></a>

## uid property — virtual_server.request_logging_profile / 0855aadc8a06 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-d254a9561b1ed441ffb009c0d97d2a176ee4a227229ab7eb4e2a5c48733dbc4e"></a>

## Next pages — virtual_server.request_logging_profile / 0855aadc8a06 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c1ca32c558bc623274b7d440f5ec362ac2a7d65650d92748a8650174060f122"></a>

## virtual_server.source_port — virtual_server.source_port / d7555b0d66e7 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.source_port

<a id="canonical-ad363b517828af74b2f7de0607ed43073ab239752d743a549b36f3fcc847d40e"></a>

Type: `"single"`. Computed.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

<a id="canonical-95f885a85c4ca0c77a91de82bd0b2b194c25e7d159cf340cd394e2ecf3dc08f8"></a>

## Direct properties — virtual_server.source_port / d7555b0d66e7 / 3

- [source_port_change](data-sources--application_profiles--reference--group-003.md#canonical-7b6d36858c1269a6ce223209362d1d1d1e82c80f9cdf8770d7fb85efa86daa8b): complete subsection reference.

- [source_port_preserve](data-sources--application_profiles--reference--group-003.md#canonical-3cd4a25912a415be465c733748e34fc75be43da7effa842218388ae9291812d1): complete subsection reference.

- [source_port_preserve_strict](data-sources--application_profiles--reference--group-003.md#canonical-5a4b90e9e59e1203043611fbb7528c60707466bd0ce9f0e305237a09320af909): complete subsection reference.

<a id="canonical-b558c8798b6068395f62de1939bdbd5c79e3c03bb2f0facbd8ae820ebd622fa1"></a>

## Next pages — virtual_server.source_port / d7555b0d66e7 / 4

- [virtual_server.source_port.source_port_change](data-sources--application_profiles--reference--group-003.md#canonical-7b6d36858c1269a6ce223209362d1d1d1e82c80f9cdf8770d7fb85efa86daa8b)
- [virtual_server.source_port.source_port_preserve](data-sources--application_profiles--reference--group-003.md#canonical-3cd4a25912a415be465c733748e34fc75be43da7effa842218388ae9291812d1)
- [virtual_server.source_port.source_port_preserve_strict](data-sources--application_profiles--reference--group-003.md#canonical-5a4b90e9e59e1203043611fbb7528c60707466bd0ce9f0e305237a09320af909)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-7b6d36858c1269a6ce223209362d1d1d1e82c80f9cdf8770d7fb85efa86daa8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-312d7b5e11a1a6310aebd10aeaefec30bfe62ce6d1395a1d04b4bdb75c44068b"></a>

## virtual_server.source_port.source_port_change — virtual_server.source_port.source_port_change / ae490038b0ba / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- virtual_server.source_port.source_port_change

<a id="canonical-4fb102546ee7433007d57646ddee19ebdd648aa79b2b3bcc20f9734116144591"></a>

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

<a id="canonical-71ea8ef01a90da10d15f54e8a72b4e79588aca2a9983cf5c975f86b493345d63"></a>

## Direct properties — virtual_server.source_port.source_port_change / ae490038b0ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2d710bfd670eab05eafbefe9ceafd757bd1bda38dd023e57070394b5e2bab6a"></a>

## Next pages — virtual_server.source_port.source_port_change / ae490038b0ba / 4

- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3cd4a25912a415be465c733748e34fc75be43da7effa842218388ae9291812d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7be0b4bdb95ccfa11f04d3936d5544df246728bc13e01085fa54af0d03f38b92"></a>

## virtual_server.source_port.source_port_preserve — virtual_server.source_port.source_port_preserve / d94776dea0f6 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- virtual_server.source_port.source_port_preserve

<a id="canonical-da91e57ce73f7e99acedeeef22892a41b34f21fc63a01bc31fe8b7bfedb4b63b"></a>

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

<a id="canonical-c1a3a6e79ffbdece7d754e2aa5c2560f0ec97b213ac04f420d310ade097ce168"></a>

## Direct properties — virtual_server.source_port.source_port_preserve / d94776dea0f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ed19588f25b86b31bcd7d14a5216461e94302f50fca48a02047b39d0a346bd1"></a>

## Next pages — virtual_server.source_port.source_port_preserve / d94776dea0f6 / 4

- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-5a4b90e9e59e1203043611fbb7528c60707466bd0ce9f0e305237a09320af909"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dfffb258312e23202a63a7f5b713a7596b7757d23d95362816937c9ccf6f8c6"></a>

## virtual_server.source_port.source_port_preserve_strict — virtual_server.source_port.source_port_preserve_strict / d039c3e8df53 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- virtual_server.source_port.source_port_preserve_strict

<a id="canonical-93b5d4c9ab521cefdb5df5d99c459a4a8c5127b0129049cfa15962003c4177d8"></a>

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

<a id="canonical-c8c50255bc02db1121317cc46f78b350dd37228cd543ec11a8fbf3dfa6572e40"></a>

## Direct properties — virtual_server.source_port.source_port_preserve_strict / d039c3e8df53 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5cd673a672af1ab1b19cb02d1c8899caaadf43cc425ecbe26120e56af269351d"></a>

## Next pages — virtual_server.source_port.source_port_preserve_strict / d039c3e8df53 / 4

- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-2c3db883e913163179c953dff01dfd2f3fdccc8020837b14abe11c6e7a6b105a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63166904b71367e5a22458c3b85071d69835b0da1ce605632e114d2ed336b84a"></a>

## virtual_server.statistics_profile — virtual_server.statistics_profile / 5027dc6ea709 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.statistics_profile

<a id="canonical-8c63ada2758f01b939e0d1e1ef09aecd8f10abb3945d84e973e82b982e1e6a87"></a>

Type: `"list"`. Computed.

Configuration parameter for statistics profile.

Upstream description:

Configuration parameter for statistics profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e453dd294e8cd0d7be7b905bc422dfc71c261cacca113e65a843fc6ecf53a19d"></a>

## Direct properties — virtual_server.statistics_profile / 5027dc6ea709 / 3

<a id="canonical-b21123a591a3a0185a49c00c3db2d1938f0fed586514246c9a5f520eef8ac7f1"></a>

<a id="canonical-afee0b0d6401389d1cf2d8b3cdf40f2a155316f3fbbda53de0e9e05904c47748"></a>

## kind property — virtual_server.statistics_profile / 5027dc6ea709 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-649caae6df455908facdf6b01b6c75fc83bc5b03dabbaa5d57ddf944a1ce9bdd"></a>

<a id="canonical-09c6de9a0ec7b5bcaf5bd02dab6c54dedaf4f91aee4d1e3f56ee4fc65662ab3e"></a>

## name property — virtual_server.statistics_profile / 5027dc6ea709 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-efc17eb45f5dfbab6c2bb1df657a670797e69834d93ce4a02818770429265fe4"></a>

<a id="canonical-cfa471f9a8ed333523e52591503c2b71738e10615cc55f93a9ff205862a79ff4"></a>

## namespace property — virtual_server.statistics_profile / 5027dc6ea709 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-e0383c1ef3c9388180333756d0428519fc42d245bf81328b32a6320e8ddfd3b1"></a>

<a id="canonical-e747e9e379ae91fa14938165bbe114979524e562de24f95120c3df6bc17e1d51"></a>

## tenant property — virtual_server.statistics_profile / 5027dc6ea709 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-d126ce4776172715ccf9e1d2159ba689e3aae18756cba9f0fcf68bbbaf31bcbc"></a>

<a id="canonical-4cf03bee3b2104026b18f115df385ffd8c9410d80a3903aaec6f3085efc0efe4"></a>

## uid property — virtual_server.statistics_profile / 5027dc6ea709 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-ee38bce5beb2a100ef27ae22ce3cc26ab5596844a5e61673514d3d49ff792591"></a>

## Next pages — virtual_server.statistics_profile / 5027dc6ea709 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0faec8322a3e82ec2297018332c812f7fe731b929187f2e44ec949ceb51a8a5"></a>

## virtual_server.tcp — virtual_server.tcp / f4308fe44aba / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.tcp

<a id="canonical-fed3b36f37e7c02ef01066cef90f506b4200d146691c163d934491714451bed0"></a>

Type: `"single"`. Computed.

TCP profiles.

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

<a id="canonical-0af8e7da007fbfcb81d262d9fae1cc88adb341d8c4057c903c8ad45ee5f1cedb"></a>

## Direct properties — virtual_server.tcp / f4308fe44aba / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3406fd434f9d32acb62b88db8de2d0501be6d33c448836a06bb1ea63e80c709a): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-a7931da6d76f83bdd5e6a9369529444fa7f3bb48c77f6ba9895ab14e3253698f): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-def953d6c9b3d4d2008e0e3bd0d72e43e393e393c3b2ca87cf327665406fecc1): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-59a5964ae7ec2fd97a57e4703a2acdfce62ae2c197629b4987550f5b2d320ad5): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-6d68f738d643a7945a88cbb2d9614ce9b9606daf09e9251d65986d558d392e5d): complete subsection reference.

<a id="canonical-7044837284371329de5731072d4a067b307ac2a7f5ce0373aa4aa4341c23bf4b"></a>

## Next pages — virtual_server.tcp / f4308fe44aba / 4

- [virtual_server.tcp.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3406fd434f9d32acb62b88db8de2d0501be6d33c448836a06bb1ea63e80c709a)
- [virtual_server.tcp.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-a7931da6d76f83bdd5e6a9369529444fa7f3bb48c77f6ba9895ab14e3253698f)
- [virtual_server.tcp.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-def953d6c9b3d4d2008e0e3bd0d72e43e393e393c3b2ca87cf327665406fecc1)
- [virtual_server.tcp.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-59a5964ae7ec2fd97a57e4703a2acdfce62ae2c197629b4987550f5b2d320ad5)
- [virtual_server.tcp.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-6d68f738d643a7945a88cbb2d9614ce9b9606daf09e9251d65986d558d392e5d)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3406fd434f9d32acb62b88db8de2d0501be6d33c448836a06bb1ea63e80c709a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfbc561fca5bec6ec749323806938526c835b8e34f058d549ad439ef8800557a"></a>

## virtual_server.tcp.client_ssl_profile — virtual_server.tcp.client_ssl_profile / f85176286e38 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- virtual_server.tcp.client_ssl_profile

<a id="canonical-6becd4ce7fc44af63bdd7bd2a5337ac39609fc2edd60384df022d7a0a0ff5489"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9b65999da1ced90d1760db51caa0ad61b607ed2691dbd4cf408ee0138c296cdb"></a>

## Direct properties — virtual_server.tcp.client_ssl_profile / f85176286e38 / 3

<a id="canonical-6811c599e8bf8411c76fe469e97ee3c50a05ec4fa4542281767b7912fe508b35"></a>

<a id="canonical-ca8c52861bc6e0e859710e389f0e663e19281bc64f16345148c4f19c1793575c"></a>

## kind property — virtual_server.tcp.client_ssl_profile / f85176286e38 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-61bbd436ad7fbe7051a82c41306382c2f6d0688cb6f82a89a7ad5b8a5ad01462"></a>

<a id="canonical-c4f626e55daf1a51903bf726391622fcabbb5d3c0e51a2631fca8a970ac3fbc4"></a>

## name property — virtual_server.tcp.client_ssl_profile / f85176286e38 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1a2785e0f2a08687dcf12d69cea302abbf1386a88f4de8043bd3f932a764eebc"></a>

<a id="canonical-127be98135fe622c7fe488ab758956696d80c01ff353640fc68e2e06b109bc40"></a>

## namespace property — virtual_server.tcp.client_ssl_profile / f85176286e38 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-461fa089a5957a838ea7178ceb6be3cdf509d19770c54f284ed9047001e93d1a"></a>

<a id="canonical-697c536f7c4fdfa970a056cd96afdf85ead18baf29ef72a943d16d02c5442261"></a>

## tenant property — virtual_server.tcp.client_ssl_profile / f85176286e38 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-991a15b602076359e1f770748d3b707d8bfe8ddf87e2b840910864f93b12fb9a"></a>

<a id="canonical-2af6d75d2cbe92dc0596f03ea9798089981e2e17b7336b6b764250375dd8d5a7"></a>

## uid property — virtual_server.tcp.client_ssl_profile / f85176286e38 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-ed2ef0f0e8ae6c85e43f7e4f17bd826b5574c54d17f6dabca246d6658ff9e3b3"></a>

## Next pages — virtual_server.tcp.client_ssl_profile / f85176286e38 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-a7931da6d76f83bdd5e6a9369529444fa7f3bb48c77f6ba9895ab14e3253698f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-781cb8eda46997c65f210aca64ff8d34c1489ea8c14d848ec1426c5fe3fbc1d4"></a>

## virtual_server.tcp.ocsp_profile — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- virtual_server.tcp.ocsp_profile

<a id="canonical-ea5e2eba3da843952d83ef1d58125020784cbe09c83b5687c6e70bb166bf588c"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-40650218fcc9234325d206d6480ab0f8a8fad06f550c377fabe195466e4e69e3"></a>

## Direct properties — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 3

<a id="canonical-9f4cd9f54b82832b309f6c38cf8fba8cc3296ce3fa39f73223980024e4e21a3e"></a>

<a id="canonical-5800149d2c40a472614ea72b415929bc34e3fa85cdb9755cdef07baaeb7b395b"></a>

## kind property — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-73c92c8c1002ec35a44af767cbe2dd502ce9f665066ec5286526f54dfe7dec3a"></a>

<a id="canonical-8962998911d18a548dd9658913c454e3b6e8532103274f7956dd4308c1a183c5"></a>

## name property — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-dcb08e05d34e14396334e1031f26405909b454fdefa3214fe6aca20da25bd3c5"></a>

<a id="canonical-8e8d47770eefcc3a54ce9c60cf302e404e18216c904e70a033fa373618c9c0db"></a>

## namespace property — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3a2731fb22e0b952d0473100a46aa47ababed01db36542e068fba7e36abcbb6f"></a>

<a id="canonical-1d77431282a76179545435e8f0b97859aa430a91745d4f2c1ab0d75f2f3f4cb8"></a>

## tenant property — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-ef7352c2aa96b88465b6492a78d95f13ed848c15cfea76ae10b9ea2b79ef2962"></a>

<a id="canonical-9b240aeca5fca3577686a1e87d1b5d0614480ad270d5036d5d1915726238b2ac"></a>

## uid property — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-a3c1166b770c0eba664c008fe65d29d91d79246af1a706d3425bbe41a95015cc"></a>

## Next pages — virtual_server.tcp.ocsp_profile / 93ad5b694d36 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-def953d6c9b3d4d2008e0e3bd0d72e43e393e393c3b2ca87cf327665406fecc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57015ddcf37ea4ab3bd70904717ca60e5d42f1d60e1d6ca63d28490cfd8888d5"></a>

## virtual_server.tcp.server_ssl_profile — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- virtual_server.tcp.server_ssl_profile

<a id="canonical-7190c5c4c134028b638b0785b8ae92f8215cc96f575a3e1d394f5720b7cb205e"></a>

Type: `"list"`. Computed.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-90b980a3d796668d10adf7c1939eb5224746bb7285108d8bf18db01baaf2163a"></a>

## Direct properties — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 3

<a id="canonical-dba3204329536526c42cb36e102676af2d46948ec422b6f3e1b6847eefb0afd0"></a>

<a id="canonical-3e9afad21c3bb18216159428bcca7ee7a32e7e532cd2fa82c411c8a5784190fc"></a>

## kind property — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-e82c9f044af52c325cd32957a3f61eb42f952eee539c68ca85b3eb171ea59a1a"></a>

<a id="canonical-56615eefcadafd04c47c5e75a3276316320cec611a62d52d89ae98338c9557bd"></a>

## name property — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-d5ad219b232cfa683ecc27ed8d124012481411ad520005c5c2f2c7e267863f7b"></a>

<a id="canonical-8748dc33a9e62f1d88b59b65c7fb2010e1cc3f0af30eebc53b3aae4380b48afd"></a>

## namespace property — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-8a523c6fafb0d46270ae4fb0f901b8285de3e9ffa2d5f5e09fbc95260e1cccfd"></a>

<a id="canonical-2383cd878bd94bfc8c895a3db4f4b8f3639561ac114689dc03826fc67abbc5a3"></a>

## tenant property — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1380e0d9b387c3c90ee1df25dd0348f3ef8966126de5fe71d2a6c4754368ad0a"></a>

<a id="canonical-896ebb2a86f7da61f4d2886ae9b19d78a129794e7e8c77cdec35f23ea51f6dcc"></a>

## uid property — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-8e95debe391a7676571a789d603081040d987fbf3379b15b8ea1b91929804e6b"></a>

## Next pages — virtual_server.tcp.server_ssl_profile / 9c474a80bf8c / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-59a5964ae7ec2fd97a57e4703a2acdfce62ae2c197629b4987550f5b2d320ad5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-642fed06b8e0c0188ab9ed9afdf6500acaeab8cbd9e4d6b33e1e2a1fa9ed682d"></a>

## virtual_server.tcp.tcp_client_profile — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- virtual_server.tcp.tcp_client_profile

<a id="canonical-2f709128592e3af6ceb0b049905d4c8b84bc61257894947d89536ef44c66e294"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-828c6a7079a8a68bd2b21fc9a419ce0770c829d4e937a9a23b89a5536676df4d"></a>

## Direct properties — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 3

<a id="canonical-cea36b67b0b9992ca66fd77a4de6e25ccbaf8dd789f59f667dd84f658e2a464b"></a>

<a id="canonical-176b25a1f7887c41518b80a0fa86a5af0f7b89e832d39c5bf6431d6dde7c88d2"></a>

## kind property — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-339cea565067082078abad2ddb0a5a5d6cd74b9bddc6677352a673278dc1df4d"></a>

<a id="canonical-38c1a68d310e5ad344252d2151e75eab9b1140623139078f66bb7a4e183ae02a"></a>

## name property — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-818389ef58313e1e4e58a54de540443cb2fe63cd9fb68044af577396e63bfd21"></a>

<a id="canonical-0a023c71b41b7fa43c44343ecff39219a137c8c52d9cc12a76750238f7ee6320"></a>

## namespace property — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-664320f73d93bf8c61cdb1ac0ca7e8e1c15dec76d13bd42aefaf41d754fbf45a"></a>

<a id="canonical-7394b218480858e33c4d3be5e749e43f5b80f5fa73cf8c30e959b100a3c95d34"></a>

## tenant property — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2fa6b3c5662d9355393deff203f45d3f572d91b54890d032692e200c281a4da4"></a>

<a id="canonical-0c28db41d1d2af9d51fc227b5a18c33774eb00bc8aca46aced3a3f19eb5f6984"></a>

## uid property — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-206d256dd098875eac3c82eb1d16dae7227a4b1cad80b88dce7caa031ec713df"></a>

## Next pages — virtual_server.tcp.tcp_client_profile / aa192b36c8cf / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-6d68f738d643a7945a88cbb2d9614ce9b9606daf09e9251d65986d558d392e5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77288d1a9ad959080acef9e3b3a44ca76c59f2bf3d4571a31f399360786a2ef4"></a>

## virtual_server.tcp.tcp_server_profile — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- virtual_server.tcp.tcp_server_profile

<a id="canonical-ab44608225a6ab07a35eaaa98010b5ee0a454d0df672599fd170fe7a969dcfe9"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-13100d2390381b4092d519b133b1cbf2adacc03e06e2e59dd0c039d52f5455ac"></a>

## Direct properties — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 3

<a id="canonical-bd0bbe56573f934579b9834b4ba824ef8925c0f631584913185142e18266c148"></a>

<a id="canonical-94f0e228e75c51dd39b7fc27df2825599577a60bee90ce01811d7a94f77dc46c"></a>

## kind property — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-83347ebd1043c69f339775ce94bdfb66c6de01778d208ea248f9564b6b6609d0"></a>

<a id="canonical-87272779e4113f0a84e29e64ff44ba5bbd3ccfa8158b9d97da0aa04fdd6aa80f"></a>

## name property — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-faa2112423a167e5586e12c02ddd13d07e54788234b1d63db9d8063292b1250b"></a>

<a id="canonical-b1c977f6753e93d2abf40fc1d57887f0dabcac58d222e4a5eb37c75381910673"></a>

## namespace property — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-396c0164845bc25d7d477c33b49184007bae206b5c353cfd6a3add209162152d"></a>

<a id="canonical-1b5da8196f21b20500d8032600bccc28d614d814ad9d4738802cf58c019daec9"></a>

## tenant property — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-d68b9e1c328b65fb3cce0050fb12b189a6129d62eb3a6b2e6d2d9dac147050d0"></a>

<a id="canonical-6eb95287ce2654145859c51041763f28b745d6ea3dda254b45d298d45add45ba"></a>

## uid property — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-553023ba3322d5b11de1c39894a42d6b4310a3f48bf32a151f3878e591ae56cd"></a>

## Next pages — virtual_server.tcp.tcp_server_profile / a195f8e12ea5 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b4a763d250e56076b91ab170572cca7b79b23eb8b3024e339de4793b9c7cf3c"></a>

## virtual_server.udp — virtual_server.udp / 11c6430f232c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.udp

<a id="canonical-ec4acd18442e0e69df2ef446188dfaacef92f9137e352e59ad0b651b89e711ae"></a>

Type: `"single"`. Computed.

UDP profiles.

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

<a id="canonical-816474226e6a35b21cd458f43c16803999e07393be555e4388d38f7d9f31c439"></a>

## Direct properties — virtual_server.udp / 11c6430f232c / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-5ca05704abf9af21af705f966a19debb8ebc2e927c6be2c4b435bc2bc1c619cd): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-6fa5e262ad70846886d13af44605b78a8260ba773b3316043268fc2a07097578): complete subsection reference.

- [udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-7f6f2c2b43983e18c5749e25d9942eaa118b2137eee7cb0e8bef1d510311b65a): complete subsection reference.

- [udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2b748d51d8dece2134bef7598615ff61264203ff33df05db6d69360aa5b33ca3): complete subsection reference.

<a id="canonical-24eb7a7650ffa11d0e4001c54fe7fa107a6425d67ce52f4c54227332ca633180"></a>

## Next pages — virtual_server.udp / 11c6430f232c / 4

- [virtual_server.udp.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-5ca05704abf9af21af705f966a19debb8ebc2e927c6be2c4b435bc2bc1c619cd)
- [virtual_server.udp.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-6fa5e262ad70846886d13af44605b78a8260ba773b3316043268fc2a07097578)
- [virtual_server.udp.udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-7f6f2c2b43983e18c5749e25d9942eaa118b2137eee7cb0e8bef1d510311b65a)
- [virtual_server.udp.udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2b748d51d8dece2134bef7598615ff61264203ff33df05db6d69360aa5b33ca3)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-5ca05704abf9af21af705f966a19debb8ebc2e927c6be2c4b435bc2bc1c619cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1098adb5d5743a32367f57ec543eb454635a65196b36541599a6e9db815d25f"></a>

## virtual_server.udp.client_ssl_profile — virtual_server.udp.client_ssl_profile / d4a901119d76 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- virtual_server.udp.client_ssl_profile

<a id="canonical-ab54d95f11fc2ee97cca03519e5ab6eaa168b9cc48e77c05447fd88decc34716"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2a97d425153a94a92725674eba15f26cb41f56a0c09c174e2cb122b8cce8cee5"></a>

## Direct properties — virtual_server.udp.client_ssl_profile / d4a901119d76 / 3

<a id="canonical-3ac464f4a5922e72a84223e9dbb59a4caafc7420abf965a45d25b755ed903b9f"></a>

<a id="canonical-e0995728ebb42164849ca7dd3caf0225b5020b416c77abac7d8b55fe0b5cf020"></a>

## kind property — virtual_server.udp.client_ssl_profile / d4a901119d76 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1e4ecef1b49112cae943af2c88afbe26269b567917a487fdb30d11209acfeb25"></a>

<a id="canonical-12e37400b42c83d872afb893495e9378d0da4c1ac68b40030d6cc2d1d80f8ec4"></a>

## name property — virtual_server.udp.client_ssl_profile / d4a901119d76 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-644ef6ec0583c430dafb65342f13d2075a0a6591668fee22798a5ec352c691c4"></a>

<a id="canonical-098d2ea3c5fa7d9725e0a3eb5a2fffcdd9792e09f2f4a7df5316ad2ba7142d37"></a>

## namespace property — virtual_server.udp.client_ssl_profile / d4a901119d76 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-25d5a383ea7155c2d103b444895cb03d9501272995623a85501cf14b2379a9b0"></a>

<a id="canonical-0a86063a682ad60d8e685d020de9640f82bc0b612fedee4e6088d2f1d077ce01"></a>

## tenant property — virtual_server.udp.client_ssl_profile / d4a901119d76 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-a5431232b819578f61cd8289689d06a886270d1e2f33fdce859d6dc32dc9b6d4"></a>

<a id="canonical-05085d41f66d0fbeaab501a01265be1b479138753deaf1fc9424107f0d4c6528"></a>

## uid property — virtual_server.udp.client_ssl_profile / d4a901119d76 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-a5c53f7a0cdf9a29341adce19df21e64e8e72bd3a87da07e14a6e65861f5c13c"></a>

## Next pages — virtual_server.udp.client_ssl_profile / d4a901119d76 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-6fa5e262ad70846886d13af44605b78a8260ba773b3316043268fc2a07097578"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ea952cc1d5aa0cc45af9077c3b2a83d975b6c642d8674fe2d64b696d636eea1"></a>

## virtual_server.udp.server_ssl_profile — virtual_server.udp.server_ssl_profile / 4a6902380515 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- virtual_server.udp.server_ssl_profile

<a id="canonical-113c6c7fbc1ec2f04c809603d5bd35cc52490f8d9d64deaaa77765bcb122e5d0"></a>

Type: `"list"`. Computed.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9d2b4aa3ee2a0554daf54b6e75bd9c5be38208f6397b35dd5c2041638874f7f3"></a>

## Direct properties — virtual_server.udp.server_ssl_profile / 4a6902380515 / 3

<a id="canonical-beec88a2b11075f5e18edff61da77733b1059650335da63c1ecdf50288e5c356"></a>

<a id="canonical-0e921142f641019b3cc1b4011ca58b061890c47df8a7c908bd90257b2fbe520f"></a>

## kind property — virtual_server.udp.server_ssl_profile / 4a6902380515 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-a9cd946b1f38f186f702b8d1fda417d9847a812c2ddf5ff37b0a2ef57118749d"></a>

<a id="canonical-0a26c853bf7de1845feb0e77baffcd224ca37b4946bcc6468f0dcd90f781058b"></a>

## name property — virtual_server.udp.server_ssl_profile / 4a6902380515 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-508d560d98eeb784b1cc1a82c620c7f767da089524611eabb117624580546145"></a>

<a id="canonical-046042cf8a0293440e1ae790a25295489155de95adcff0cbfdaa720c9e86f7a3"></a>

## namespace property — virtual_server.udp.server_ssl_profile / 4a6902380515 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-6b698e52808ff6c37a68bda6c03456d0e6df4b24b6bb024987b4d22e7224ad04"></a>

<a id="canonical-773ddc8b24a8c94ee3890b9ff1b852c01670e8e3bf8a0ac5b2a2e28c23ae8bad"></a>

## tenant property — virtual_server.udp.server_ssl_profile / 4a6902380515 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-ad72e40adb053a0d8f9b4934802eae5d00cb27c7a6e984fd246bd016121514d3"></a>

<a id="canonical-cdc9561fddc3d81af304a442affd7cab04f4576ab2cf5f79188191ff81397583"></a>

## uid property — virtual_server.udp.server_ssl_profile / 4a6902380515 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-cdff57fb90562a8896838a94f71a62f4b328cef62400b90b02b06058d6186fde"></a>

## Next pages — virtual_server.udp.server_ssl_profile / 4a6902380515 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-7f6f2c2b43983e18c5749e25d9942eaa118b2137eee7cb0e8bef1d510311b65a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3193229e0873a519e6fad4c0a812de515991b980fdb60dd9b79a73a795494de1"></a>

## virtual_server.udp.udp_client_profile — virtual_server.udp.udp_client_profile / 61b529cee03c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- virtual_server.udp.udp_client_profile

<a id="canonical-4a287d25dd2f51853eac7bfc346f70b2c33101f37d4d142b0275987bab0093d0"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f856a9d0ede3f1f87ba72a20d0e5d93913b4c3e942fccbc56d3149272c70d3f0"></a>

## Direct properties — virtual_server.udp.udp_client_profile / 61b529cee03c / 3

<a id="canonical-fde67c0e8239a86bdea73fbed61c54aeb34094c964604ce4e33c022f3b9bf08f"></a>

<a id="canonical-ac124972ca476246371182b562661a643785905c1e79dfbbd02a0817c306c9dc"></a>

## kind property — virtual_server.udp.udp_client_profile / 61b529cee03c / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-78afd2af2586cbd82b0a7420e4620abc9fa017628d91ec37aebce715f874fa22"></a>

<a id="canonical-e1025e8e95eb5f5afba53f99ccdcbeb0625258f73fd904ea64e5e98d714a8fab"></a>

## name property — virtual_server.udp.udp_client_profile / 61b529cee03c / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-b367ea08a04834bae4bf87c6c934a3b99c4b4ca5f819a7ce6778288df9dbcbdd"></a>

<a id="canonical-da72ce11e34cd5021a81358f567ea7bbfdbb341df5a780fd63226643cc231d86"></a>

## namespace property — virtual_server.udp.udp_client_profile / 61b529cee03c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-55dc3f2796ce12cc3ecc8f7a2905aa1cf5d4cd4fed73c9caf139069afb8338eb"></a>

<a id="canonical-2e94fcf73551d9b9d2402a9df64da588a40bbaf1876f18b410c16aebd60a4d7b"></a>

## tenant property — virtual_server.udp.udp_client_profile / 61b529cee03c / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-dbee2710dea23677e4d9c68675867c61a63d2d5a8055a2ee513b89d9ee1c846f"></a>

<a id="canonical-9ad6eee41f8d599c2f7247eb857bc9cd3c5b2e206940c52dc2004137de6dba53"></a>

## uid property — virtual_server.udp.udp_client_profile / 61b529cee03c / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-870ae002a943d35fed263ebbe8497ba8ab9639ffe2c5b6ae659bab331f3b4478"></a>

## Next pages — virtual_server.udp.udp_client_profile / 61b529cee03c / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-2b748d51d8dece2134bef7598615ff61264203ff33df05db6d69360aa5b33ca3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ef363050df618182b17b1e58bdd7bb92b9f7168e2400c92f58f803bdcd146fb"></a>

## virtual_server.udp.udp_server_profile — virtual_server.udp.udp_server_profile / e3acc3742767 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- virtual_server.udp.udp_server_profile

<a id="canonical-6e92e9fae03d12e44a42d646f5ef3e621c3827c8b10da3ffc912cee3cec7a35f"></a>

Type: `"list"`. Computed.

Configuration parameter for udp server profile.

Upstream description:

Configuration parameter for udp server profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c62bfeada9dd2520363da9ab4d4c6fce6757e267551e05b4908a1bd35c17ef3f"></a>

## Direct properties — virtual_server.udp.udp_server_profile / e3acc3742767 / 3

<a id="canonical-31ba8ff8262f4af9af428d96b0f6b13385ae283f5f651afb9b60f7982d7865ba"></a>

<a id="canonical-47f06f11ca7ddeca2cb7ec7fb76354a11d375b0ce4bfa8e2faf55a6b7f99c3c2"></a>

## kind property — virtual_server.udp.udp_server_profile / e3acc3742767 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-715d9c3e85718431be75de5c94927bb9fdd6ab1fb922b4594f87cfd1f56fb977"></a>

<a id="canonical-95e5de1e2cf3bd98bd3a93363028ae1aa971a94ba8c8255a5e32e9d5f1fbb90f"></a>

## name property — virtual_server.udp.udp_server_profile / e3acc3742767 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-50e6cd84cb7a6202fe774e173df025ba5e27adc665f392bb22a3eec870dcf1fb"></a>

<a id="canonical-b160b20c72b12184f2296714f175c646c6eb1c5d7681ddfd9b67e76d4f78abe6"></a>

## namespace property — virtual_server.udp.udp_server_profile / e3acc3742767 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-08bfd6347868017170671eff6fbd78b8eb195e222dbaa38f938c0ab5466fff99"></a>

<a id="canonical-ee03ae94d0dcbc32e858348ecdfb7262bf77e62c851a15f823291711fdeaec3e"></a>

## tenant property — virtual_server.udp.udp_server_profile / e3acc3742767 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-f27b48285694654d1bb1b9759d1b5919f4e5c621a16e71645bf9110e6a8a5d2e"></a>

<a id="canonical-46a53f14783e9b38a649b06d3064118606e92efd5a0c6478a515b4d82d2f9b7f"></a>

## uid property — virtual_server.udp.udp_server_profile / e3acc3742767 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-e573f82985e37ea6d553ee4aa2e87cd5885f9652c17e369b11312b4fc3fa5e93"></a>

## Next pages — virtual_server.udp.udp_server_profile / e3acc3742767 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f3ec03e09d1005c092e7bc4feb5543e6c38f0d803dad1552087d805d013173e"></a>

## virtual_server.virtual_server_state — virtual_server.virtual_server_state / d6dd660f0d7d / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.virtual_server_state

<a id="canonical-89b19c385bb0df18a103d7cdde9054f96f70b81b27afcc611e2ae38962137b90"></a>

Type: `"single"`. Computed.

Displays the current state on the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

<a id="canonical-3b5b26d48f0599f494dbfc7020edb82ddf6f1f21f964279aeab6c2dd64dace70"></a>

## Direct properties — virtual_server.virtual_server_state / d6dd660f0d7d / 3

- [state_disabled](data-sources--application_profiles--reference--group-003.md#canonical-a6ba90e672f3b5f5fffaa7e118293dc0f9471815c2350a52fb4418df15019460): complete subsection reference.

- [state_enabled](data-sources--application_profiles--reference--group-003.md#canonical-e3a7db0e9e219628a0110262cceed299a40ab237defeff93f980e6c5bc543cbb): complete subsection reference.

<a id="canonical-d3c6f210add62af8ac7a9f765956d5797af4c17b0241600994dee026f62ddd26"></a>

## Next pages — virtual_server.virtual_server_state / d6dd660f0d7d / 4

- [virtual_server.virtual_server_state.state_disabled](data-sources--application_profiles--reference--group-003.md#canonical-a6ba90e672f3b5f5fffaa7e118293dc0f9471815c2350a52fb4418df15019460)
- [virtual_server.virtual_server_state.state_enabled](data-sources--application_profiles--reference--group-003.md#canonical-e3a7db0e9e219628a0110262cceed299a40ab237defeff93f980e6c5bc543cbb)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-a6ba90e672f3b5f5fffaa7e118293dc0f9471815c2350a52fb4418df15019460"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16c6d661268a302f3555cad30b5b005bfe2b94e98ffb0f566b3fef49e568c3ce"></a>

## virtual_server.virtual_server_state.state_disabled — virtual_server.virtual_server_state.state_disabled / 4d37fec7021f / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c)
- virtual_server.virtual_server_state.state_disabled

<a id="canonical-bfb0fbd34a39406765d93e7ea7128ff07984685a2a47757d91a083beb53af121"></a>

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

<a id="canonical-3b3f4caa74207822f146d52381f8a7c1abbad23364cdaa76d0d9e76b8245b22b"></a>

## Direct properties — virtual_server.virtual_server_state.state_disabled / 4d37fec7021f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06b3fd88b5f052844fd9aff662e613289c45153d99f6359260f0066a315cc538"></a>

## Next pages — virtual_server.virtual_server_state.state_disabled / 4d37fec7021f / 4

- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-e3a7db0e9e219628a0110262cceed299a40ab237defeff93f980e6c5bc543cbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34ef5ed7b2e494c5dac952ec3db495d890945c50dbebfd3efb6113c606f0db08"></a>

## virtual_server.virtual_server_state.state_enabled — virtual_server.virtual_server_state.state_enabled / b6b9fbf5a6dc / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c)
- virtual_server.virtual_server_state.state_enabled

<a id="canonical-9713f599f7c135068f02155ddd77a98abe3da7c62a31f2b43d4462c0527cd32c"></a>

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

<a id="canonical-0c6c8d80be827d77e5a8c7b888b51d2e1806c3e5fc997c667643bbce5e605756"></a>

## Direct properties — virtual_server.virtual_server_state.state_enabled / b6b9fbf5a6dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b4f9b8c4536c1c034079c860b07db66fa55c82184c11c9ea7ad6c767ed912709"></a>

## Next pages — virtual_server.virtual_server_state.state_enabled / b6b9fbf5a6dc / 4

- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
