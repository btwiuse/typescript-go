//// [tests/cases/compiler/declFileModuleAssignmentInObjectLiteralProperty.ts] ////

//// [declFileModuleAssignmentInObjectLiteralProperty.ts]
module m1 {
    export class c {
    }
}
var d = {
    m1: { m: m1 },
    m2: { c: m1.c },
};

//// [declFileModuleAssignmentInObjectLiteralProperty.js]
var m1;
(function (m1) {
    class c {
    }
    m1.c = c;
})(m1 || (m1 = {}));
var d = {
    m1: { m: m1 },
    m2: { c: m1.c },
};
