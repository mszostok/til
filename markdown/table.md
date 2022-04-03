## Diff tables

You can use HTML tables to show the diff quite nicely, see:

- New feature: **Support multiple input parameters on Implementation**
     <table>
     <tr>
     <td> Before </td> <td> After </td>
     </tr>
     <tr>
     <td>

     ```yaml
     inject:
       additionalInput:
         additional-parameters:
             replicaCount: 3
     ```

     </td>
     <td>

     ```yaml
     inject:
       additionalParameters:
         - name: rds-parameters
           value:
             replicaCount: 3
     ```

     </td>
     </tr>
     </table>
