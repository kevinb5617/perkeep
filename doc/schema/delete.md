# Delete Claim

A claim can delete a permanode or another claim.

(Un)Deletions are not considered as modifications, so the claimDate of a delete
claim is never considered as a modtime in the context of time constrained
searches.

Example:

    {"camliVersion": 1,
     "camliType": "claim",
     "camliSigner": "sha1-f2b0b7da718b97ce8c31591d8ed4645c777f3ef4",
     "claimDate": "2010-07-10T17:20:03.9212Z",
     "claimType": "delete",
     "target": "sha1-ab6dacb972eeee72df2a846aab7d751b5856a1a0", // the permanode or claim being deleted.
     <REQUIRED-JSON-SIGNATURE>
    }
