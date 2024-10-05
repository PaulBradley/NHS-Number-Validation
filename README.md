# NHS Number Validation

Whist working on a LIMS data migration project, I needed a function to validate patient NHS Numbers. All [NHS Numbers](https://www.datadictionary.nhs.uk/attributes/nhs_number.html) have a check digit which validates all the previous nine numbers within the 10 digit number.

I'd already written the `isValidNHSNumber` function in Go for another project, but the LIMS data migration was being developed using C# and .NET 8. I didn't want to have to rewrite the function, so this repository takes the Go implementation and allows it to be compiled to a Windows DLL file.

The Windows DLL file can then be imported into a C# program and reused.

The `isValidNHSNumber` function accepts an NHS number as a string and returns and integer. The return value of `0` indicates the validation check has failed. A return value of `1` denotes the validation check passed. The function accepts a string because sometimes NHS numbers are represented with spaces to make it easier to read. The function strips out all space characters before performing the validation.


## Compile the Windows DLL

To compile the DLL run the following `make` command:

```bash
make dll
```

## Using the function in C#

```csharp
// import the DLL and define
// the function signature
[System.Runtime.InteropServices.DllImport(@"nhs-number-validation.dll")]
static extern int isValidNHSNumber(string value);

Console.WriteLine(isValidNHSNumber("999 9999 994"));   // returns false
Console.WriteLine(isValidNHSNumber("999 9999 999"));   // returns true
Console.WriteLine(isValidNHSNumber("111 111   1111")); // returns true
```