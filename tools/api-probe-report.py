from pathlib import Path
import json, os
private = Path(os.environ['XCSH_ACC_PRIVATE'])
patterns = {
 **{f"name_limit_{value}": [str(value)] for value in [16,20,24,30,31,32,36,40,48,50,52,54,56,60,63,64]},
    'length_limit': ['length','longer','characters','too long'],
    'size_limit': ['size','bytes','too large'],
    'name_field': ['name'],
    'label_field': ['label'],
    'name_32_limit': ['32'],
    'name_48_limit': ['48'],
    'name_63_limit': ['63'],
    'name_64_limit': ['64'],
    'max_limit': ['maximum','max','limit'],
    'minimum_limit': ['minimum','min','at least'],
    'reserved_identity': ['reserved','forbidden prefix'],
    'definition_field': ['definition'],
    'resource_field': ['resource'],
    'unique_identity': ['unique','duplicate'],
    'oneof_constraint': ['oneof','one of','exactly one'],
    'missing_input': ['missing','empty','no value'],
    'not_found': ['not found','does not exist'],
    'create_operation': ['create','creation'],
    'metadata_rejected': ['metadata','annotations','labels'],
    'invalid_parameters': ['invalid request parameters'],
    'not_allowed_namespace': ['not allowed in namespace','not supported in namespace','shared namespace','namespace shared','invalid namespace'],
    'name_rejected': ['invalid name','name is','name must','name too','name length'],
    'request_rejected': ['bad request','invalid request','invalid argument','validation failed'],
    'already_exists': ['already exists','conflict'],
    'permission_missing': ['permission','not authorized','forbidden'],
    'empty_definition': ['no api','no swagger','must provide','at least one'],
    'pattern_constraint': ['pattern','regex','must match','does not match'],
    'namespace_constraint': ['namespace'],
    'url_scheme_constraint': ['url','uri','scheme'],
    'document_fetch_failed': ['fetch','download','not found','no such'],
    'document_parse_failed': ['parse','unmarshal','format','json','yaml'],
    'missing_required_input': ['required','must specify','mandatory','empty'],
    'unsupported_input': ['unsupported','not supported','not allowed'],
    'identity_constraint': ['name','label','length','size'],
    'swagger_reference': ['swagger','openapi','specification'],
    'authorization_denied': ['permission','denied','unauthorized'],
}
result=[]
for variant in ['empty','relative','absolute']:
    raw=json.loads((private/f'api-probe-{variant}.json').read_text())
    if raw.get('variant') != variant: raise ValueError('diagnostic variant mismatch')
    accepted=raw.get('accepted') is True
    if not accepted and not isinstance(raw.get('error'),str): raise ValueError('diagnostic result missing')
    error=raw.get('error','').lower()
    result.append({'variant':variant,'accepted':accepted,'signatures':[name for name,terms in patterns.items() if any(term in error for term in terms)]})
Path('test-reports/api-probe.json').write_text(json.dumps({'version':1,'variants':result},indent=2)+'\n')
