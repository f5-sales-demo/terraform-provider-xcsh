---
page_title: "xcsh_cloud_credentials reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials reference."
---

# xcsh_cloud_credentials reference

<a id="canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302113020030201-2113113221222331-1032121013102110-2203331111120132-0323213003110303-3013000321110002-1122123221123300-1300020202020032"></a>

## Property reference — Property reference / 200301030012 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- Property reference

<a id="canonical-3301323120203002-1212230301020210-0323111022311033-0201010001002223-3033022112201002-3331102232200330-2200022332312310-0032203303102202"></a>

## Direct properties — Property reference / 200301030012 / 3

<a id="canonical-2233030303112110-1113133003333220-0303220222112230-1332110221202133-1320211333322113-3301133122122012-2223323211312313-3102100202211033"></a>

<a id="canonical-1233212232301011-3300122331222120-3333000030101211-3030230101310122-3030132120032201-3233302201113103-0322123113232203-3001332310113303"></a>

## annotations property — Property reference / 200301030012 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
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

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303): complete subsection reference.

- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213): complete subsection reference.

- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203): complete subsection reference.

- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232): complete subsection reference.

<a id="canonical-1101031231300233-3333312013123233-0021113011202210-2121032320123110-2100310010131310-3202303100321213-2202112020300022-1002320232003203"></a>

<a id="canonical-0330232012120022-1312011010223322-1100200131023031-2222023121101303-3213321323220210-2122301321133001-2002032323122201-3320033123300332"></a>

## description property — Property reference / 200301030012 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220): complete subsection reference.

<a id="canonical-2322023113103001-1232323021101321-3122333020012022-0023031200310331-0000102303100012-3033111131300032-3131011222320000-2122123102132032"></a>

<a id="canonical-0023232001000311-2332333000130122-1222233123103303-2012120022121010-1122021110300001-2031103013011231-1132303133013031-0231010021110332"></a>

## ID property — Property reference / 200301030012 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0030130310122212-2112022002033023-2121102200300111-1321300323121230-2102222011322132-0101221220220120-3223221010313331-3210311201211213"></a>

<a id="canonical-2130200300121301-0223301111313100-2313312231021211-1130321302110111-0313233322122033-3221010312323203-2233113123100102-1211013102021211"></a>

## labels property — Property reference / 200301030012 / 7

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

<a id="canonical-3232200220232023-1333231200200213-2112103202303330-1200210213110321-3003200303301023-0111202130022213-0230310101330111-3211210030033233"></a>

<a id="canonical-1232311000230100-0131100101220200-1301022303203330-0121031021132020-3333133101330220-3102210103110300-0310301213031300-2323131000222100"></a>

## name property — Property reference / 200301030012 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1220003322023030-2123010032203102-0201133231012023-2311112201202230-0132320012130000-3233121113321332-3223301011000123-0322310130011012"></a>

<a id="canonical-1011030311223131-0302201032032020-3002222020013332-1032012102322231-1023121210131321-0012311321203033-2320023032003212-2113330023231220"></a>

## namespace property — Property reference / 200301030012 / 9

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0200322012222101-0300223203223210-0113023133121202-1012020200321032-0103112123012001-1311213020331312-0313231212220133-2201120300223022"></a>

## All schema paths — Property reference / 200301030012 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_credentials--reference--group-001.md#canonical-2233030303112110-1113133003333220-0303220222112230-1332110221202133-1320211333322113-3301133122122012-2223323211312313-3102100202211033) |
| `aws_assume_role` | [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3233023331021113-0132033102010323-0333230210210222-3332222311123311-1311321013323013-2221330313003333-0113023133000130-0130211331211231) |
| `aws_assume_role.custom_external_id` | [aws_assume_role.custom_external_id](data-sources--cloud_credentials--reference--group-001.md#canonical-1320322003130230-0010301001223302-3001201110322010-3031113323023002-0102121011122031-2001303301030312-0121110011321133-1222212213211332) |
| `aws_assume_role.duration_seconds` | [aws_assume_role.duration_seconds](data-sources--cloud_credentials--reference--group-001.md#canonical-3232300231022133-3122120031300303-0220333022320331-1111300032123300-2112013101032231-3300222103133122-3210330212033301-2222001222020023) |
| `aws_assume_role.external_id_is_optional` | [aws_assume_role.external_id_is_optional](data-sources--cloud_credentials--reference--group-001.md#canonical-2033220100223012-3332221202321120-0323310132020203-0232003123311011-3020210303031033-0012212322313300-1022013212110312-1021000302301231) |
| `aws_assume_role.external_id_is_tenant_id` | [aws_assume_role.external_id_is_tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-2103323023213033-3201321132112233-0113311012210211-3110122021033220-0221013212221333-3130301121330230-0331233203030002-1121310001001122) |
| `aws_assume_role.role_arn` | [aws_assume_role.role_arn](data-sources--cloud_credentials--reference--group-001.md#canonical-3232322113132322-2303310102100112-0303220103301103-1033011123122331-2000131111023000-1122302012232022-1103221001232113-3201102101300100) |
| `aws_assume_role.session_name` | [aws_assume_role.session_name](data-sources--cloud_credentials--reference--group-001.md#canonical-1220113233121111-3002021320013121-1111012331122220-0113032033000301-3223032121201302-1103123212220022-2302212333223122-0113312300321103) |
| `aws_assume_role.session_tags` | [aws_assume_role.session_tags](data-sources--cloud_credentials--reference--group-001.md#canonical-3103223021112321-3210220011310301-2131030003022222-2222021101332012-1111302001203221-0231102211222311-3120001101133022-2232111130021011) |
| `aws_secret_key` | [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-3233330000131200-2002101002210021-1122223232221333-0221133021300122-1032030020012030-0003022102232012-1201322330112133-1322022132123213) |
| `aws_secret_key.access_key` | [aws_secret_key.access_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0023001031211130-3110133013110303-3213133321300113-3202113223112130-2113322021332031-0203302302333323-2203012332210101-0121000220030012) |
| `aws_secret_key.secret_key` | [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-1101232101323033-0203232201032033-0112322230332031-2211321320123203-0113211320000201-3111223232123031-0330031032130001-2032031200103132) |
| `aws_secret_key.secret_key.blindfold_secret_info` | [aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-3231131311131122-3331022030220100-3100222301100222-0320130320333000-3021210030231223-0120202101031112-2032123333333102-3313032121312031) |
| `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-0231332010032313-1312213123210010-0203320111332230-2103310201301232-1300021000332113-1332202323312331-1020003311313002-0021023023012122) |
| `aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_secret_key.secret_key.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-2121012011320300-1222311230033101-2001333010033203-2212020222222210-1113302002002300-2203322013230123-1333232130223203-2222222200021212) |
| `aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_secret_key.secret_key.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-1113213231301302-0323313232023202-0022110303322303-0033200102012203-2202222003230103-3311112232200012-0123222210033230-1122303212012113) |
| `aws_secret_key.secret_key.clear_secret_info` | [aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-0131310211100100-1212320300011302-2021032213030221-3111001101111103-3232013002131130-1022021332311103-0010211120202211-0011203203310123) |
| `aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_secret_key.secret_key.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-0101200221200200-3113120333212201-1331011233310203-2222113021232000-2023033022313321-0231303312101313-1232222113113023-2233133100110202) |
| `aws_secret_key.secret_key.clear_secret_info.url` | [aws_secret_key.secret_key.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-3032033200020303-1231330311012133-0211230232123301-3221231322320011-0331312010313210-3023120122013033-2121231331213310-2223302213202211) |
| `azure_client_secret` | [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0003011031130323-2310222230110002-1111203122100031-1113232312221211-1313301213110020-1000011300021200-1012321231203122-0330222030132012) |
| `azure_client_secret.client_id` | [azure_client_secret.client_id](data-sources--cloud_credentials--reference--group-001.md#canonical-2321103023331200-2201130230323212-0301022000003222-0301322213000201-0032230322231103-1312102203220301-0203200123032020-3231012202230021) |
| `azure_client_secret.client_secret` | [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-3122023113313312-3323020133032331-1201310200000300-3110202010220211-2102130012101310-1030102321331023-1213032010330202-1000323201011311) |
| `azure_client_secret.client_secret.blindfold_secret_info` | [azure_client_secret.client_secret.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-2211011101201321-0201103223223300-3000113032310010-0023213022020230-1233322010132100-3232333332330200-1211231201233113-3133332013301223) |
| `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` | [azure_client_secret.client_secret.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-3222112303313221-3032313001013103-1310312222223213-3222313302321013-2210113033223121-2221211011200213-0333101002300312-3101122100120113) |
| `azure_client_secret.client_secret.blindfold_secret_info.location` | [azure_client_secret.client_secret.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-2321133132300321-3100133231203030-2200303100122322-0033301211301220-2233123003211222-0031202323112101-3231230220321123-3313310111301303) |
| `azure_client_secret.client_secret.blindfold_secret_info.store_provider` | [azure_client_secret.client_secret.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-2022322021103131-3130321122033223-0131230032023230-1123231011230332-1111212222212033-1203031001123111-2233020331223111-2000000322102310) |
| `azure_client_secret.client_secret.clear_secret_info` | [azure_client_secret.client_secret.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-3321200311320111-3332200023101300-3122222032002132-0132213211312213-0030313022230300-1101120320030022-1321202312101002-1100303022001021) |
| `azure_client_secret.client_secret.clear_secret_info.provider_ref` | [azure_client_secret.client_secret.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-3312320322011221-2121323132032300-0311223122033201-2023221313000302-3303331333223130-3020231311032123-1113323000231210-1322000320230023) |
| `azure_client_secret.client_secret.clear_secret_info.url` | [azure_client_secret.client_secret.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-1132320303313302-1300200011113301-1110032302022222-0002313320101232-0300313330220211-2113130330001112-3001002200021021-1211101222322203) |
| `azure_client_secret.subscription_id` | [azure_client_secret.subscription_id](data-sources--cloud_credentials--reference--group-001.md#canonical-2123112123000013-3300010310020211-1231121110322033-1023301101023302-1113033331200010-3001310200001001-0120010223223111-1332031202133232) |
| `azure_client_secret.tenant_id` | [azure_client_secret.tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-2223330201222020-0001010113020131-1233202200000012-0311031031223101-0022323112121120-0132231030003310-3122131130033110-0100331001000231) |
| `azure_pfx_certificate` | [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-0120122132310303-2033211130110121-2001012233223321-2213310101321111-1320012222120032-2111223001132201-1211303123213131-2100332011022222) |
| `azure_pfx_certificate.certificate_url` | [azure_pfx_certificate.certificate_url](data-sources--cloud_credentials--reference--group-001.md#canonical-2000100222020223-1003322121120003-1133311011232302-0100312231121100-2021003013002030-3323231023011031-2120221312112222-2332133021213033) |
| `azure_pfx_certificate.client_id` | [azure_pfx_certificate.client_id](data-sources--cloud_credentials--reference--group-001.md#canonical-0230131113001022-3013003030231000-3233333213031201-3301002333333213-0320110132223300-2123102322101002-2102001232001111-3310010300010303) |
| `azure_pfx_certificate.password` | [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-0110000213031313-0003330231222222-0312201302323100-0003103330203213-3122102131333213-0021010210001111-0303222131331300-1123321113013201) |
| `azure_pfx_certificate.password.blindfold_secret_info` | [azure_pfx_certificate.password.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-0220020232322301-1310012103011230-1000230133031131-0303013203021111-1222221122023322-0111032220300330-3103202100032032-1201112111000303) |
| `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` | [azure_pfx_certificate.password.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-3113111303320223-1311213123221332-3330001312012120-1121022230000211-1210301013100222-0313231000100130-0323213320333101-0231231110103010) |
| `azure_pfx_certificate.password.blindfold_secret_info.location` | [azure_pfx_certificate.password.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-3120013200030310-2333320031003103-2213101103013211-2312121311101301-1103302101232312-0133031332310201-0313222101202300-0222323310130321) |
| `azure_pfx_certificate.password.blindfold_secret_info.store_provider` | [azure_pfx_certificate.password.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-3132113022133122-2322120233323010-2231132022013020-3123023002212002-2011113230021122-2203223110121210-0213321211120223-1231203013103310) |
| `azure_pfx_certificate.password.clear_secret_info` | [azure_pfx_certificate.password.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1121330220202021-2220333301333322-1101322212021333-1002130123203123-0231200013123132-0220220320201121-3013102110212020-1230120100122232) |
| `azure_pfx_certificate.password.clear_secret_info.provider_ref` | [azure_pfx_certificate.password.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-3320023310320023-2132302111312302-0201321302301311-1311101022312322-1320330301011101-3000330003233120-3222321121222231-0123231330132200) |
| `azure_pfx_certificate.password.clear_secret_info.url` | [azure_pfx_certificate.password.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-3130200301110220-2130221000210103-2300312122130331-0031120221233000-0301221322332322-3033302002112302-0320223303010010-2330333012013100) |
| `azure_pfx_certificate.subscription_id` | [azure_pfx_certificate.subscription_id](data-sources--cloud_credentials--reference--group-001.md#canonical-3132131022222213-1111223121330013-1233300131312220-1203200013200223-3000032303131130-3332002332223010-3120222012123201-3131000003113111) |
| `azure_pfx_certificate.tenant_id` | [azure_pfx_certificate.tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-1300212103301032-2232211112312022-0211021131020213-3310111223202110-0001300213231320-0103212302322333-3203323312120200-0233332120020132) |
| `description` | [description](data-sources--cloud_credentials--reference--group-001.md#canonical-1101031231300233-3333312013123233-0021113011202210-2121032320123110-2100310010131310-3202303100321213-2202112020300022-1002320232003203) |
| `gcp_cred_file` | [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-0013303022211201-0113013303033133-0103212211330211-0033222331302200-1000103213213223-2031101233123223-2301132011102131-2321000032232302) |
| `gcp_cred_file.credential_file` | [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-0300111030211300-3032110222123120-3033022200312133-2122311030313333-2133233211102133-1022102203302111-0201233223000332-1213011010022301) |
| `gcp_cred_file.credential_file.blindfold_secret_info` | [gcp_cred_file.credential_file.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-0033023213311331-3110301123101231-2011201302111030-0031020302023310-1123123122113221-3331113010111302-2221222120201011-1132303331113212) |
| `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-1202202133322331-0103120103222311-0331032020130020-2301233001031302-2200003101213310-1011113321003220-3133032212212312-3311033230210232) |
| `gcp_cred_file.credential_file.blindfold_secret_info.location` | [gcp_cred_file.credential_file.blindfold_secret_info.location](data-sources--cloud_credentials--reference--group-001.md#canonical-3312102323031012-3110001213123121-3123132131220310-3312323122310023-2221003010310011-2111002210331012-1130333121330133-3031310201322122) |
| `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.store_provider](data-sources--cloud_credentials--reference--group-001.md#canonical-2320323102311230-3331301020300110-1223010031310230-3101320021323121-2033332310302322-1331131101123232-2121300332301232-3032122122301301) |
| `gcp_cred_file.credential_file.clear_secret_info` | [gcp_cred_file.credential_file.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-2001330120000203-1332020203100200-0202202123232213-2023120320201331-3323231301332311-0231112313313211-2311300303122320-0130132023333310) |
| `gcp_cred_file.credential_file.clear_secret_info.provider_ref` | [gcp_cred_file.credential_file.clear_secret_info.provider_ref](data-sources--cloud_credentials--reference--group-001.md#canonical-0013301331133013-0011023003321300-1120312013013231-1201301011222331-1111002302031132-1133033030013101-2312031322010121-3032200103320333) |
| `gcp_cred_file.credential_file.clear_secret_info.url` | [gcp_cred_file.credential_file.clear_secret_info.url](data-sources--cloud_credentials--reference--group-001.md#canonical-1002031212130211-0211332100002032-0013113021132300-1303232012022321-2221120021323311-3211222313330333-2231300330103313-2322010312010110) |
| `id` | [ID](data-sources--cloud_credentials--reference--group-001.md#canonical-2322023113103001-1232323021101321-3122333020012022-0023031200310331-0000102303100012-3033111131300032-3131011222320000-2122123102132032) |
| `labels` | [labels](data-sources--cloud_credentials--reference--group-001.md#canonical-0030130310122212-2112022002033023-2121102200300111-1321300323121230-2102222011322132-0101221220220120-3223221010313331-3210311201211213) |
| `name` | [name](data-sources--cloud_credentials--reference--group-001.md#canonical-3232200220232023-1333231200200213-2112103202303330-1200210213110321-3003200303301023-0111202130022213-0230310101330111-3211210030033233) |
| `namespace` | [namespace](data-sources--cloud_credentials--reference--group-001.md#canonical-1220003322023030-2123010032203102-0201133231012023-2311112201202230-0132320012130000-3233121113321332-3223301011000123-0322310130011012) |

<a id="canonical-0322122312001333-2021303022320231-2323101131102103-0233131233131031-3310212131101223-0302003033102031-1322330031230122-2003001002001231"></a>

## Next pages — Property reference / 200301030012 / 11

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333130022322023-3203231003312311-1122210332100230-2231231113222200-1010211311331133-3112133103110212-0121313100332031-2102110112301130"></a>

## aws_assume_role — aws_assume_role / 130333323232 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- aws_assume_role

<a id="canonical-3233023331021113-0132033102010323-0333230210210222-3332222311123311-1311321013323013-2221330313003333-0113023133000130-0130211331211231"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_assume\_role, aws\_secret\_key, Azure\_client\_secret, Azure\_pfx\_certificate,
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

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3233023331021113-0132033102010323-0333230210210222-3332222311123311-1311321013323013-2221330313003333-0113023133000130-0130211331211231)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-3233330000131200-2002101002210021-1122223232221333-0221133021300122-1032030020012030-0003022102232012-1201322330112133-1322022132123213)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0003011031130323-2310222230110002-1111203122100031-1113232312221211-1313301213110020-1000011300021200-1012321231203122-0330222030132012)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-0120122132310303-2033211130110121-2001012233223321-2213310101321111-1320012222120032-2111223001132201-1211303123213131-2100332011022222)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-0013303022211201-0113013303033133-0103212211330211-0033222331302200-1000103213213223-2031101233123223-2301132011102131-2321000032232302)

Select alternatives according to the provider validators above.

<a id="canonical-3200321012322000-2030221112102010-3301322010111112-3102133213033111-1313000311331032-1222130002112110-3223022113302303-1301211303110003"></a>

## Direct properties — aws_assume_role / 130333323232 / 3

<a id="canonical-1320322003130230-0010301001223302-3001201110322010-3031113323023002-0102121011122031-2001303301030312-0121110011321133-1222212213211332"></a>

<a id="canonical-0023121132113320-3321023330100210-0002033022310130-3303012333312313-0210211330202122-2012220232322030-0110313201130223-3112233313003231"></a>

## custom_external_id property — aws_assume_role / 130333323232 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3232300231022133-3122120031300303-0220333022320331-1111300032123300-2112013101032231-3300222103133122-3210330212033301-2222001222020023"></a>

<a id="canonical-2030001301000231-0032032332021111-1033112302221123-3213122020011301-0122120331111313-3132302221220113-3021003022123211-2221123231013103"></a>

## duration_seconds property — aws_assume_role / 130333323232 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [external_id_is_optional](data-sources--cloud_credentials--reference--group-001.md#canonical-1111312120100210-0223030012102003-2002211112110022-1223030110102223-1113200123003201-1323023301021000-1202030011112320-2131312112312103): complete subsection reference.

- [external_id_is_tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-3112020101022232-3101230110212113-2321111223231203-0102022101223230-0110230133332101-0011212322003132-2311223021031332-0100020223132301): complete subsection reference.

<a id="canonical-3232322113132322-2303310102100112-0303220103301103-1033011123122331-2000131111023000-1122302012232022-1103221001232113-3201102101300100"></a>

<a id="canonical-2111001003102332-0333200220323320-3100300120312300-2220010030222232-0011233021013201-3030001303311033-0033120333210112-0001020300321320"></a>

## role_arn property — aws_assume_role / 130333323232 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1220113233121111-3002021320013121-1111012331122220-0113032033000301-3223032121201302-1103123212220022-2302212333223122-0113312300321103"></a>

<a id="canonical-1133313103122011-1222332112120301-3120233013203032-2133111320231333-1211111312331330-2322211031030112-1333010332212220-3000232003222232"></a>

## session_name property — aws_assume_role / 130333323232 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3103223021112321-3210220011310301-2131030003022222-2222021101332012-1111302001203221-0231102211222311-3120001101133022-2232111130021011"></a>

<a id="canonical-1323333331213121-3123323122232300-3220003021320213-3003311003332131-0000303132202210-0330000021302113-2221330031032301-3020133121102001"></a>

## session_tags property — aws_assume_role / 130333323232 / 8

Type: `["map", "string"]`. Computed.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 40,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
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

<a id="canonical-3210101021310133-1102012323022211-0132002010030301-2103000320001103-0201112332303000-1313230021010321-0120131022103033-3001101022211213"></a>

## Next pages — aws_assume_role / 130333323232 / 9

- [aws_assume_role.external_id_is_optional](data-sources--cloud_credentials--reference--group-001.md#canonical-1111312120100210-0223030012102003-2002211112110022-1223030110102223-1113200123003201-1323023301021000-1202030011112320-2131312112312103)
- [aws_assume_role.external_id_is_tenant_id](data-sources--cloud_credentials--reference--group-001.md#canonical-3112020101022232-3101230110212113-2321111223231203-0102022101223230-0110230133332101-0011212322003132-2311223021031332-0100020223132301)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1111312120100210-0223030012102003-2002211112110022-1223030110102223-1113200123003201-1323023301021000-1202030011112320-2131312112312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111303321200321-2011312120303330-2020002301333222-0133311013200333-3201030200001222-2300130123022223-0213211020303130-2100013023210001"></a>

## aws_assume_role.external_id_is_optional — external_id_is_optional / 123320030300 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303)
- aws_assume_role.external_id_is_optional

<a id="canonical-2033220100223012-3332221202321120-0323310132020203-0232003123311011-3020210303031033-0012212322313300-1022013212110312-1021000302301231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external ID is optional.

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

<a id="canonical-1122213223010333-2123132230310320-2203313323333010-2302011011120302-1022210232003022-2330301003121231-3200323301013021-3032111303031311"></a>

## Direct properties — external_id_is_optional / 123320030300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212022001023212-3311201001211123-3110320112030233-3212101321130202-2022230312101020-3303033121003230-3311032123321030-1230220123012120"></a>

## Next pages — external_id_is_optional / 123320030300 / 4

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-3112020101022232-3101230110212113-2321111223231203-0102022101223230-0110230133332101-0011212322003132-2311223021031332-0100020223132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202330133112213-3000020310023001-1201023102322300-2323111330030222-2220002223022301-0233130100002221-0223031003122323-1000111123013223"></a>

## aws_assume_role.external_id_is_tenant_id — external_id_is_tenant_id / 002210202020 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303)
- aws_assume_role.external_id_is_tenant_id

<a id="canonical-2103323023213033-3201321132112233-0113311012210211-3110122021033220-0221013212221333-3130301121330230-0331233203030002-1121310001001122"></a>

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

<a id="canonical-2132220300003023-3013212301333123-1322002120013302-1133212213221310-2001203101320302-2331223120121000-3013132233311131-2130031002032301"></a>

## Direct properties — external_id_is_tenant_id / 002210202020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201002032111001-2012213022102013-0122321102010310-2231103132310002-1103310330022221-0301111301310013-2220212232310212-0112021121223303"></a>

## Next pages — external_id_is_tenant_id / 002210202020 / 4

- [aws_assume_role](data-sources--cloud_credentials--reference--group-001.md#canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131101110021221-2122122011132311-2310302022021131-3101003112103120-2223110123000333-2112131010330210-3322331131222121-2122130320033033"></a>

## aws_secret_key — aws_secret_key / 032320323232 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- aws_secret_key

<a id="canonical-3233330000131200-2002101002210021-1122223232221333-0221133021300122-1032030020012030-0003022102232012-1201322330112133-1322022132123213"></a>

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

<a id="canonical-3101102322021103-2122220102121100-0332011123200331-2213321031012302-3303323210322200-1112132231232233-0232132112131212-0320003010120122"></a>

## Direct properties — aws_secret_key / 032320323232 / 3

<a id="canonical-0023001031211130-3110133013110303-3213133321300113-3202113223112130-2113322021332031-0203302302333323-2203012332210101-0121000220030012"></a>

<a id="canonical-1211112322103021-1031012113011112-2020211123202110-1100021003010331-0123112211010220-0300122000333030-1303220333322230-0321331332332123"></a>

## access_key property — aws_secret_key / 032320323232 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121): complete subsection reference.

<a id="canonical-1032200130012031-0310313210111103-1210303031120111-3021121033112320-3210201132021220-2330223023212130-3002102310323320-2032130330312213"></a>

## Next pages — aws_secret_key / 032320323232 / 5

- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120011023223202-3010020122201031-0230320133222201-1030211201131022-3300302010301202-0122333320330322-0332030211332132-2301123100203133"></a>

## aws_secret_key.secret_key — secret_key / 331110333003 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213)
- aws_secret_key.secret_key

<a id="canonical-1101232101323033-0203232201032033-0112322230332031-2211321320123203-0113211320000201-3111223232123031-0330031032130001-2032031200103132"></a>

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

<a id="canonical-3112232222332100-2301032122022221-2020303330220332-1003001212001300-0211102203200022-3211230121022212-0001132321330012-3102030030133031"></a>

## Direct properties — secret_key / 331110333003 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-2303230322213013-3131202032001010-2001032223311232-0301121022220020-1300122001103030-2313233221230302-1322313302021021-2211323122203132): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1003301322220213-0203002023322022-2031123111220002-3222221100100201-2011302333213220-1312320003223003-1122022123201023-1002021211333102): complete subsection reference.

<a id="canonical-3302300030031000-3033110211133213-2011230320123122-3231003002130032-1102131202012210-1320010120121202-0230311022213212-0100012012301322"></a>

## Next pages — secret_key / 331110333003 / 4

- [aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-2303230322213013-3131202032001010-2001032223311232-0301121022220020-1300122001103030-2313233221230302-1322313302021021-2211323122203132)
- [aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1003301322220213-0203002023322022-2031123111220002-3222221100100201-2011302333213220-1312320003223003-1122022123201023-1002021211333102)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-2303230322213013-3131202032001010-2001032223311232-0301121022220020-1300122001103030-2313233221230302-1322313302021021-2211323122203132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201133213302133-2231101230303222-2033230101020210-1322123330202012-0001311220312311-3220003112201031-3033321023101122-3011233231100112"></a>

## aws_secret_key.secret_key.blindfold_secret_info — blindfold_secret_info / 233112220020 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213)
- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121)
- aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-3231131311131122-3331022030220100-3100222301100222-0320130320333000-3021210030231223-0120202101031112-2032123333333102-3313032121312031"></a>

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

<a id="canonical-2001222323010302-2031020021011320-3302022322001002-2031213111112300-3231000232001102-0113100221133321-0213033303232022-0202011200033211"></a>

## Direct properties — blindfold_secret_info / 233112220020 / 3

<a id="canonical-0231332010032313-1312213123210010-0203320111332230-2103310201301232-1300021000332113-1332202323312331-1020003311313002-0021023023012122"></a>

<a id="canonical-0203213132220003-2122223112323020-2321232313003211-3102311100110031-1323323320101231-3002030030213222-0110331320221200-2031331320031000"></a>

## decryption_provider property — blindfold_secret_info / 233112220020 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2121012011320300-1222311230033101-2001333010033203-2212020222222210-1113302002002300-2203322013230123-1333232130223203-2222222200021212"></a>

<a id="canonical-0023311101002023-3102220102312201-1131130302201323-3321220121221233-1100330221133200-3133212003230120-2120212123123131-1302203032202203"></a>

## location property — blindfold_secret_info / 233112220020 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1113213231301302-0323313232023202-0022110303322303-0033200102012203-2202222003230103-3311112232200012-0123222210033230-1122303212012113"></a>

<a id="canonical-2111110032222210-0202322233003123-1020320102120303-1121333221233332-0211203020003012-3133100200301130-1323002130001332-2222221201210302"></a>

## store_provider property — blindfold_secret_info / 233112220020 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0310232122011202-3003001211112222-3032300002212102-2222123000103321-1211121323011131-0300330230303310-0112012212312221-2100220120301333"></a>

## Next pages — blindfold_secret_info / 233112220020 / 7

- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1003301322220213-0203002023322022-2031123111220002-3222221100100201-2011302333213220-1312320003223003-1122022123201023-1002021211333102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121013002012020-3003222131001231-0023202112330131-3230201333232210-2113332011203031-1321130220232312-0121020303011212-2322011121122202"></a>

## aws_secret_key.secret_key.clear_secret_info — clear_secret_info / 302133323213 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [aws_secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-2102012002200222-1332023020013113-3310012010032313-0312331320223212-0320002011230002-3101303122122120-1121003223011022-0030321111323213)
- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121)
- aws_secret_key.secret_key.clear_secret_info

<a id="canonical-0131310211100100-1212320300011302-2021032213030221-3111001101111103-3232013002131130-1022021332311103-0010211120202211-0011203203310123"></a>

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

<a id="canonical-1111330210313103-0030020133130021-3302120020222203-0132001033333313-3030323002312131-0121323032132123-2101112001232332-0311230222121101"></a>

## Direct properties — clear_secret_info / 302133323213 / 3

<a id="canonical-0101200221200200-3113120333212201-1331011233310203-2222113021232000-2023033022313321-0231303312101313-1232222113113023-2233133100110202"></a>

<a id="canonical-0213013311300333-3133201300021231-1032323210013231-0223022130223111-3323033113032121-3320110113303011-2210131020300000-2313103110313223"></a>

## provider_ref property — clear_secret_info / 302133323213 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3032033200020303-1231330311012133-0211230232123301-3221231322320011-0331312010313210-3023120122013033-2121231331213310-2223302213202211"></a>

<a id="canonical-3021003131111302-0031313011013112-1012320332113122-2310033131032122-3033033022200310-0132012001123030-3202233212110120-2321032231233001"></a>

## URL property — clear_secret_info / 302133323213 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2333201013103223-1231121013133222-3101102132231213-1120301313122000-2212310310303220-2220223032220231-0021110022230023-2031030302311003"></a>

## Next pages — clear_secret_info / 302133323213 / 6

- [aws_secret_key.secret_key](data-sources--cloud_credentials--reference--group-001.md#canonical-0000030301313211-1102212031230300-0232233312132322-0132120333302323-2323012023231132-2120330131130311-0323222220323032-1110131002131121)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132021020101033-0321313300102311-2112100133311112-3131102013322320-0123033200230321-2232122100310322-1322323331121310-3201311012203121"></a>

## azure_client_secret — azure_client_secret / 031032023101 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- azure_client_secret

<a id="canonical-0003011031130323-2310222230110002-1111203122100031-1113232312221211-1313301213110020-1000011300021200-1012321231203122-0330222030132012"></a>

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

<a id="canonical-1313023210131231-3111321200331233-2330120323303221-2131020230033312-3130032232233203-3332110321331113-1101011101213222-0303121202122000"></a>

## Direct properties — azure_client_secret / 031032023101 / 3

<a id="canonical-2321103023331200-2201130230323212-0301022000003222-0301322213000201-0032230322231103-1312102203220301-0203200123032020-3231012202230021"></a>

<a id="canonical-2333231131123023-2210033313233021-1031031023202213-1232003030122320-0201311031120332-0102112321213122-3033310022023113-2033012232100110"></a>

## client_id property — azure_client_secret / 031032023101 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100): complete subsection reference.

<a id="canonical-2123112123000013-3300010310020211-1231121110322033-1023301101023302-1113033331200010-3001310200001001-0120010223223111-1332031202133232"></a>

<a id="canonical-3311011230001200-3010332212211330-3323300313010312-0202030110323103-0302202323200120-2001122330032121-2333230322310210-0102230303001101"></a>

## subscription_id property — azure_client_secret / 031032023101 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2223330201222020-0001010113020131-1233202200000012-0311031031223101-0022323112121120-0132231030003310-3122131130033110-0100331001000231"></a>

<a id="canonical-1213211202033101-0133310022103013-0203311113002230-2022203331012033-0120111133333021-2211003323000212-1023201321113130-1111113033311021"></a>

## tenant_id property — azure_client_secret / 031032023101 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2331120320312222-1112031113010321-0122323031123021-0232003323123233-1113200223331203-2131022013120021-2333021213020222-0002222012332323"></a>

## Next pages — azure_client_secret / 031032023101 / 7

- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103221013331322-3210211020320301-2203233112322322-1331001323030201-3123132020022030-0032221213112200-3310302311030322-2203031201322220"></a>

## azure_client_secret.client_secret — client_secret / 101332131211 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203)
- azure_client_secret.client_secret

<a id="canonical-3122023113313312-3323020133032331-1201310200000300-3110202010220211-2102130012101310-1030102321331023-1213032010330202-1000323201011311"></a>

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

<a id="canonical-1312030311011110-1202221222131101-2203223301312012-0332010103022001-0311111012333323-3011213103330311-3221033132303320-1203212220210222"></a>

## Direct properties — client_secret / 101332131211 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1310001121331112-3103330212001032-0312133021010310-3033013110030320-1313322010001323-2310133003202131-3333130102300201-3122221233331013): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1331221233331312-1112201120023032-3000221310220230-2233013110330211-0020012130200233-2313032100121030-3310002223100010-3110331031332201): complete subsection reference.

<a id="canonical-3013311033203200-3000110330132322-0302013010201133-3100101101210122-2222320001121212-2232203213111100-3110230012112010-0111233122202202"></a>

## Next pages — client_secret / 101332131211 / 4

- [azure_client_secret.client_secret.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1310001121331112-3103330212001032-0312133021010310-3033013110030320-1313322010001323-2310133003202131-3333130102300201-3122221233331013)
- [azure_client_secret.client_secret.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1331221233331312-1112201120023032-3000221310220230-2233013110330211-0020012130200233-2313032100121030-3310002223100010-3110331031332201)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1310001121331112-3103330212001032-0312133021010310-3033013110030320-1313322010001323-2310133003202131-3333130102300201-3122221233331013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022112130131013-2222301231001023-2010321202233103-2213100121310311-1220033213313020-3221312002202300-1211111101013211-1212031120220013"></a>

## azure_client_secret.client_secret.blindfold_secret_info — blindfold_secret_info / 303232321021 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203)
- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100)
- azure_client_secret.client_secret.blindfold_secret_info

<a id="canonical-2211011101201321-0201103223223300-3000113032310010-0023213022020230-1233322010132100-3232333332330200-1211231201233113-3133332013301223"></a>

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

<a id="canonical-1301331312103021-3302313130220110-0231301200312313-1130223112220313-2121213331203001-3220013110222010-2030030110003231-1002013031333210"></a>

## Direct properties — blindfold_secret_info / 303232321021 / 3

<a id="canonical-3222112303313221-3032313001013103-1310312222223213-3222313302321013-2210113033223121-2221211011200213-0333101002300312-3101122100120113"></a>

<a id="canonical-0023302230130321-0203133001023022-1023231311021322-3023201011113100-2010212202231003-1211220303102000-1331023222030222-3301333120003130"></a>

## decryption_provider property — blindfold_secret_info / 303232321021 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2321133132300321-3100133231203030-2200303100122322-0033301211301220-2233123003211222-0031202323112101-3231230220321123-3313310111301303"></a>

<a id="canonical-1013311300022002-1202022132332010-2010011330100120-0302303132101313-0122311102310032-2302233310132120-2201001020210122-0212221102210011"></a>

## location property — blindfold_secret_info / 303232321021 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2022322021103131-3130321122033223-0131230032023230-1123231011230332-1111212222212033-1203031001123111-2233020331223111-2000000322102310"></a>

<a id="canonical-1203023133323200-3220031200112302-0230312331233100-2120100132031121-3021102330331313-2122210303321210-2021131202012320-0011022203013310"></a>

## store_provider property — blindfold_secret_info / 303232321021 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230131210003311-1033220033012320-3012002031202221-2320112310201133-0000310022301311-3102221000310122-1123331031230102-1132332213103201"></a>

## Next pages — blindfold_secret_info / 303232321021 / 7

- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1331221233331312-1112201120023032-3000221310220230-2233013110330211-0020012130200233-2313032100121030-3310002223100010-3110331031332201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012103223131112-0132323023023312-1110133120223111-0123011331322111-3120231310002123-1112321222331302-1210310211110213-0110203012313322"></a>

## azure_client_secret.client_secret.clear_secret_info — clear_secret_info / 312321230323 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [azure_client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203)
- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100)
- azure_client_secret.client_secret.clear_secret_info

<a id="canonical-3321200311320111-3332200023101300-3122222032002132-0132213211312213-0030313022230300-1101120320030022-1321202312101002-1100303022001021"></a>

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

<a id="canonical-1301230232020012-2220321231213001-2312100121330110-3000223201013100-0113110023101100-0130323332203221-0130103103330031-2310101312132011"></a>

## Direct properties — clear_secret_info / 312321230323 / 3

<a id="canonical-3312320322011221-2121323132032300-0311223122033201-2023221313000302-3303331333223130-3020231311032123-1113323000231210-1322000320230023"></a>

<a id="canonical-0311131010101113-3121323301220122-2020120101220002-3030002032011230-2121200210100231-2221122221202100-3321010011032020-3113120220332213"></a>

## provider_ref property — clear_secret_info / 312321230323 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1132320303313302-1300200011113301-1110032302022222-0002313320101232-0300313330220211-2113130330001112-3001002200021021-1211101222322203"></a>

<a id="canonical-1103330120011322-3330203201301012-0000332031300111-1330302010232320-2212103131312131-1102321302003122-3323232112330221-1001213121111111"></a>

## URL property — clear_secret_info / 312321230323 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1320030200012321-3233032101331213-0330313213310231-1133310132320220-1103322133030232-2222231312213001-0130223320333211-0303012201111202"></a>

## Next pages — clear_secret_info / 312321230323 / 6

- [azure_client_secret.client_secret](data-sources--cloud_credentials--reference--group-001.md#canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111222313210231-2011101110133031-1303011111120331-1312310230230310-3122131012033123-3313023333023020-0012113010322311-0310031320202030"></a>

## azure_pfx_certificate — azure_pfx_certificate / 221201301210 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- azure_pfx_certificate

<a id="canonical-0120122132310303-2033211130110121-2001012233223321-2213310101321111-1320012222120032-2111223001132201-1211303123213131-2100332011022222"></a>

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

<a id="canonical-3013002313333112-2100232021011000-0011013030122003-1123211001210301-1303323122100100-1131211301321102-0203323033031112-1223301301223310"></a>

## Direct properties — azure_pfx_certificate / 221201301210 / 3

<a id="canonical-2000100222020223-1003322121120003-1133311011232302-0100312231121100-2021003013002030-3323231023011031-2120221312112222-2332133021213033"></a>

<a id="canonical-3222201223311211-3300020232310012-2011030020200220-3110321320010213-1333022100330022-0110101222112330-2312023103331333-2113012321201011"></a>

## certificate_url property — azure_pfx_certificate / 221201301210 / 4

Type: `"string"`. Computed.

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;base64 of certificate&gt;
format. Here &lt;base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Upstream description:

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;base64 of certificate&gt;
format. Here &lt;base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230131113001022-3013003030231000-3233333213031201-3301002333333213-0320110132223300-2123102322101002-2102001232001111-3310010300010303"></a>

<a id="canonical-3210310222002131-3033322130001221-1323201310100303-1222303201001313-0010010233232231-2021002232010231-0113101311111331-2032330030101112"></a>

## client_id property — azure_pfx_certificate / 221201301210 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [password](data-sources--cloud_credentials--reference--group-001.md#canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310): complete subsection reference.

<a id="canonical-3132131022222213-1111223121330013-1233300131312220-1203200013200223-3000032303131130-3332002332223010-3120222012123201-3131000003113111"></a>

<a id="canonical-2301012111033333-0133220312230212-1001011121121021-2232310132102202-0213031321211233-3330111302213220-2230221011031320-0101222231032303"></a>

## subscription_id property — azure_pfx_certificate / 221201301210 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1300212103301032-2232211112312022-0211021131020213-3310111223202110-0001300213231320-0103212302322333-3203323312120200-0233332120020132"></a>

<a id="canonical-0231202220332333-1222222233311031-2223021000131111-2110302012300101-0133312330313303-1113103011003011-3222312022202330-2302103333133332"></a>

## tenant_id property — azure_pfx_certificate / 221201301210 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2120320232130311-2110200200200210-3312311200203330-1020100333333311-3220213323322220-1201132002223303-1003211333121202-1132223220322013"></a>

## Next pages — azure_pfx_certificate / 221201301210 / 8

- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330001200020000-0220012311021321-0033331033030222-2202332221101333-1131123032001330-0120101112200121-1313101302120313-1221030333022230"></a>

## azure_pfx_certificate.password — password / 213330213102 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232)
- azure_pfx_certificate.password

<a id="canonical-0110000213031313-0003330231222222-0312201302323100-0003103330203213-3122102131333213-0021010210001111-0303222131331300-1123321113013201"></a>

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

<a id="canonical-3230212211322331-2300120313000330-3030023032301223-0103130212000122-1232300020123030-3212022031102211-1303032203220202-2111220331301032"></a>

## Direct properties — password / 213330213102 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1212200030110223-2201123122310132-3102120313332010-2301222033013013-3301200312123223-3222332020031311-3023111221322021-1333103133032020): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1201122001333213-2300323330321203-0123302211331200-2220221322111113-3232002223333130-1010233333222230-1020111022310032-3021031330332132): complete subsection reference.

<a id="canonical-0002330133002021-3113033101332230-3222311303323313-0220021113233130-3210202011010002-1310331321231301-3103122101033310-3211120223021323"></a>

## Next pages — password / 213330213102 / 4

- [azure_pfx_certificate.password.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1212200030110223-2201123122310132-3102120313332010-2301222033013013-3301200312123223-3222332020031311-3023111221322021-1333103133032020)
- [azure_pfx_certificate.password.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-1201122001333213-2300323330321203-0123302211331200-2220221322111113-3232002223333130-1010233333222230-1020111022310032-3021031330332132)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1212200030110223-2201123122310132-3102120313332010-2301222033013013-3301200312123223-3222332020031311-3023111221322021-1333103133032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021132122211033-2121211003323223-0210013011033122-2112021002222232-0030113311321020-2111321311011013-0110101211233222-1313303210132211"></a>

## azure_pfx_certificate.password.blindfold_secret_info — blindfold_secret_info / 201221222022 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232)
- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310)
- azure_pfx_certificate.password.blindfold_secret_info

<a id="canonical-0220020232322301-1310012103011230-1000230133031131-0303013203021111-1222221122023322-0111032220300330-3103202100032032-1201112111000303"></a>

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

<a id="canonical-0300331003223132-0313321310233331-0321230013013111-0302020202103220-0210333132021001-3112211112303112-0220100020011000-3233000300113223"></a>

## Direct properties — blindfold_secret_info / 201221222022 / 3

<a id="canonical-3113111303320223-1311213123221332-3330001312012120-1121022230000211-1210301013100222-0313231000100130-0323213320333101-0231231110103010"></a>

<a id="canonical-2330010233211021-0123030110000220-0113211033131103-2033331022312321-1132221312320231-2100023201123231-2013331231303102-1321332312120133"></a>

## decryption_provider property — blindfold_secret_info / 201221222022 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3120013200030310-2333320031003103-2213101103013211-2312121311101301-1103302101232312-0133031332310201-0313222101202300-0222323310130321"></a>

<a id="canonical-1033123133120100-0333100330110210-0112001201131103-0200003030200210-0323002021022011-3211222330310133-3101000313120332-0013213102021233"></a>

## location property — blindfold_secret_info / 201221222022 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132113022133122-2322120233323010-2231132022013020-3123023002212002-2011113230021122-2203223110121210-0213321211120223-1231203013103310"></a>

<a id="canonical-2031232232201320-2131111012023101-0010023222330300-1333303222133031-0230102010010300-2212111113122233-2200112313221303-0013100102112013"></a>

## store_provider property — blindfold_secret_info / 201221222022 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2200312333202122-2123002013313311-2201103312220202-3201330101111002-3300333330022031-1221221311131212-0302220221133000-1301233110102303"></a>

## Next pages — blindfold_secret_info / 201221222022 / 7

- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-1201122001333213-2300323330321203-0123302211331200-2220221322111113-3232002223333130-1010233333222230-1020111022310032-3021031330332132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132313033223103-0010213113122100-1010013033230002-3011100023300310-3322303033221211-3123210331230030-0011010000300101-0210102101330100"></a>

## azure_pfx_certificate.password.clear_secret_info — clear_secret_info / 123313013120 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [azure_pfx_certificate](data-sources--cloud_credentials--reference--group-001.md#canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232)
- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310)
- azure_pfx_certificate.password.clear_secret_info

<a id="canonical-1121330220202021-2220333301333322-1101322212021333-1002130123203123-0231200013123132-0220220320201121-3013102110212020-1230120100122232"></a>

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

<a id="canonical-3222121211200310-1200023113221203-2033210031003023-1120120111130202-2313330013011331-3021100221231212-0223110220133203-1002122312133302"></a>

## Direct properties — clear_secret_info / 123313013120 / 3

<a id="canonical-3320023310320023-2132302111312302-0201321302301311-1311101022312322-1320330301011101-3000330003233120-3222321121222231-0123231330132200"></a>

<a id="canonical-3112120001303323-3211321001101320-0131303233002000-3203032100103312-3333021201333003-0010203100101330-1111303013300123-2232231310130321"></a>

## provider_ref property — clear_secret_info / 123313013120 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3130200301110220-2130221000210103-2300312122130331-0031120221233000-0301221322332322-3033302002112302-0320223303010010-2330333012013100"></a>

<a id="canonical-2231301300331200-0133201122232231-2020003103231213-3223232121230303-0023032013301321-1210323213132203-1202313032021130-1123313013303201"></a>

## URL property — clear_secret_info / 123313013120 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1310121230332322-0232131300032130-2310102033001222-0122122122333320-2021023022120003-1202222021330213-1202000020202103-3200320013110002"></a>

## Next pages — clear_secret_info / 123313013120 / 6

- [azure_pfx_certificate.password](data-sources--cloud_credentials--reference--group-001.md#canonical-1203212212031313-1311000121012312-3233332111333130-0133110132301013-0130021033132310-0132013100121001-2201100330301301-0132022231111310)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010001023113132-2332011102332031-0203202333313230-2313110312123201-1320110202233131-1221131333002331-1313110102231331-3102212123031313"></a>

## gcp_cred_file — gcp_cred_file / 012210213223 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- gcp_cred_file

<a id="canonical-0013303022211201-0113013303033133-0103212211330211-0033222331302200-1000103213213223-2031101233123223-2301132011102131-2321000032232302"></a>

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

<a id="canonical-0111320312231013-3301122100003002-1122121330310211-3122212111320011-2120301030102231-1211201330321212-1101112321303221-2112110331103232"></a>

## Direct properties — gcp_cred_file / 012210213223 / 3

- [credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201): complete subsection reference.

<a id="canonical-3133300320111302-1123312001020210-0331020323102023-0010031233312310-2102203231113122-0303121322133001-2102231112300221-0031222203232032"></a>

## Next pages — gcp_cred_file / 012210213223 / 4

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223020121231010-3312023311000210-2101221203330202-2001133013012101-1000132100323221-0220032300331323-1222210332010312-3003033030200130"></a>

## gcp_cred_file.credential_file — credential_file / 330332301130 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220)
- gcp_cred_file.credential_file

<a id="canonical-0300111030211300-3032110222123120-3033022200312133-2122311030313333-2133233211102133-1022102203302111-0201233223000332-1213011010022301"></a>

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

<a id="canonical-1130022311001032-1020013020312000-0130232013132200-1303201320033220-2120310303100213-0002211121023103-2111011020221200-1130102033012113"></a>

## Direct properties — credential_file / 330332301130 / 3

- [blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-3030221311212313-3330012333112113-3013322321133220-1110122200110300-3310131020223311-1102102112033320-1202230001233020-3131130330012020): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-3333121211320012-3022222023100321-0123130331331031-0130111310330302-3023201233331131-0233301013103303-1112010111210313-1021101110300122): complete subsection reference.

<a id="canonical-0331230231002012-2200332333012012-3312112021013023-3030100112031002-3302200120233102-0103110113022110-1331031333231111-3223233232020010"></a>

## Next pages — credential_file / 330332301130 / 4

- [gcp_cred_file.credential_file.blindfold_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-3030221311212313-3330012333112113-3013322321133220-1110122200110300-3310131020223311-1102102112033320-1202230001233020-3131130330012020)
- [gcp_cred_file.credential_file.clear_secret_info](data-sources--cloud_credentials--reference--group-001.md#canonical-3333121211320012-3022222023100321-0123130331331031-0130111310330302-3023201233331131-0233301013103303-1112010111210313-1021101110300122)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-3030221311212313-3330012333112113-3013322321133220-1110122200110300-3310131020223311-1102102112033320-1202230001233020-3131130330012020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313303320003100-2011211112122323-0233130002001121-1200012111220232-2322131301020301-1212000023311311-0213311132321222-2220331100132132"></a>

## gcp_cred_file.credential_file.blindfold_secret_info — blindfold_secret_info / 121112202100 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220)
- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201)
- gcp_cred_file.credential_file.blindfold_secret_info

<a id="canonical-0033023213311331-3110301123101231-2011201302111030-0031020302023310-1123123122113221-3331113010111302-2221222120201011-1132303331113212"></a>

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

<a id="canonical-1200232301122013-2120123131212230-3100130300003133-2221312003023311-2033123312131001-1321211130113330-1302113333001332-0110112133110232"></a>

## Direct properties — blindfold_secret_info / 121112202100 / 3

<a id="canonical-1202202133322331-0103120103222311-0331032020130020-2301233001031302-2200003101213310-1011113321003220-3133032212212312-3311033230210232"></a>

<a id="canonical-3123301122203303-3210010022013231-1013313022222001-3313022110230111-0101201023223320-3313200312213323-1310100223001201-0200220222210223"></a>

## decryption_provider property — blindfold_secret_info / 121112202100 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3312102323031012-3110001213123121-3123132131220310-3312323122310023-2221003010310011-2111002210331012-1130333121330133-3031310201322122"></a>

<a id="canonical-2121311313023103-2313210033323331-2022011101311103-2232022301001032-0203230010212320-3212331000322323-0110130322032232-1231233021120311"></a>

## location property — blindfold_secret_info / 121112202100 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2320323102311230-3331301020300110-1223010031310230-3101320021323121-2033332310302322-1331131101123232-2121300332301232-3032122122301301"></a>

<a id="canonical-0321222321120233-1012123232032213-2330131221130132-1312211113133032-1103333003033232-1311103120323020-3102203022231201-0122331122301303"></a>

## store_provider property — blindfold_secret_info / 121112202100 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2123330102100331-0310211031011322-0203230322110132-2131313132123002-2020121321301011-3200032110332201-2212231220130301-0031222031202220"></a>

## Next pages — blindfold_secret_info / 121112202100 / 7

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)

<a id="canonical-3333121211320012-3022222023100321-0123130331331031-0130111310330302-3023201233331131-0233301013103303-1112010111210313-1021101110300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131323132113213-2222221121021132-3101303202333132-0231210310220120-1202231000331312-2301013031101003-0232310002020233-0332323311202221"></a>

## gcp_cred_file.credential_file.clear_secret_info — clear_secret_info / 020221333002 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
- [Property reference](data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [gcp_cred_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220)
- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201)
- gcp_cred_file.credential_file.clear_secret_info

<a id="canonical-2001330120000203-1332020203100200-0202202123232213-2023120320201331-3323231301332311-0231112313313211-2311300303122320-0130132023333310"></a>

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

<a id="canonical-1322233131330211-1031230103111333-2003002221200131-3033110320130221-0232113131112113-3202303211001122-0223020113002321-2222033230002233"></a>

## Direct properties — clear_secret_info / 020221333002 / 3

<a id="canonical-0013301331133013-0011023003321300-1120312013013231-1201301011222331-1111002302031132-1133033030013101-2312031322010121-3032200103320333"></a>

<a id="canonical-3210333202020330-0002330120321121-2222210311130320-2123033023313003-2100313210311030-0110303101033012-0122012133033112-2110032233133203"></a>

## provider_ref property — clear_secret_info / 020221333002 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1002031212130211-0211332100002032-0013113021132300-1303232012022321-2221120021323311-3211222313330333-2231300330103313-2322010312010110"></a>

<a id="canonical-0020121122222002-3003221133032113-0222300310222311-0130123321032021-1320323313011021-1233020203021131-0321333031202200-3320102203021132"></a>

## URL property — clear_secret_info / 020221333002 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1022010323330023-1221211001120030-0222120002032312-1213021301320021-2203222221302322-1212100223322302-3131302211133012-0103110221013033"></a>

## Next pages — clear_secret_info / 020221333002 / 6

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--reference--group-001.md#canonical-2210320220000231-3331031000103311-3330201120020032-3302312311123020-1200230310011323-2132320323231103-0231313320002033-2301203011020201)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md#canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313)
