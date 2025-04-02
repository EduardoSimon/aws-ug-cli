import json
import re


def is_valid_domain(domain):
    # Regular expression for basic domain validation
    pattern = r'^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$'
    return bool(re.match(pattern, domain))


def lambda_handler(event, context):
    # Check if domain exists in the event
    if 'domain' not in event:
        return {
            'statusCode': 400,
            'body': json.dumps('Error: domain parameter is required')
        }

    domain = event['domain']

    # Validate domain format
    if not is_valid_domain(domain):
        return {
            'statusCode': 400,
            'body': json.dumps(f'Error: Invalid domain format - {domain}')
        }

    print(f"Processing request for domain: {domain}")
    return {
        'statusCode': 200,
        'body': json.dumps(f'Cache flushed successfully for domain {domain}')
    }
