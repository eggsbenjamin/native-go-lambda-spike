### Limits

- Step Function Input : 32768 chars
- Lambda Synchronous Invocation Input : 6MB
- Lambda Asynchronous Invocation Input : 128KB

### Test Data

largest_non_compressed_stepfn.json - 32KB
largest_non_compressed_stepfn.json.gz - 607B

massive.json.10k.gz - original size : 1MB
massive.json.20k.gz - original size : 1.9MB
massive.json.30k.gz - original size : 2.8MB
massive.json.40k.gz - original size : 3.6MB
massive.json.50k.gz - original size : 4.5MB

## Notes

- it seems to be faster to execute the lambda (based on largest_non_compressed_stepfn.json) when it's gzipped and the unmarshaling/marshaling has been optimised.
