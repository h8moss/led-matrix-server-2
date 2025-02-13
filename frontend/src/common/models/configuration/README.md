# Configuration

The configuration object must be backwards compatible, so that if you upgrade the vector, you can use the same function.
For that purpose, we are adding a description of each version of each configuration:

## Version 01
```
{
    defaults: {
        module-name: { an object with the values of each form }
    }
}
```
That's right! Version 1 only keeps track of the defaults. Let's hope we don't add any more versions

