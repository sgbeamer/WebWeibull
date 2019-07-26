function printlist(input) {
    console.log(input.toString());
}

function sort(times, suspensions) {
    var sorted = [[], []];
    for(var i = 0; i < times.length; i++) {
        sorted[0][i] = times[i];
        if(suspensions[i]) {
            sorted[1][i] = 1;
        } else {
            sorted[1][i] = 0;
        }
    }
    for(var i = 0; i < sorted[0].length-1; i++) {
        var temp = sorted[0][i];
        var ind = i;
        for(var j = i+1; j < sorted[0].length; j++) {
            if(sorted[0][j] < sorted[0][ind]) {
                temp = sorted[0][j];
                ind = j;
            }
        }
        sorted[0][ind] = sorted[0][i];
        sorted[0][i] = temp;
        var btemp = sorted[1][i];
        sorted[1][i] = sorted[1][ind];
        sorted[1][ind] = btemp;
    }
    return sorted;
}
function genRevr(times) {
    var revr = [];
    for(var i = 0; i < times.length; i++) {
        revr[i] = times.length-i;
    }
    return revr;
}
function rank(adjranks, n) {
    //console.log("start rank");//
    //console.log(adjranks.length);//
    var ranks = [];
    //double n = adjranks.length;
    for(var ind = 0; ind < adjranks.length; ind++) {
        var i = adjranks[ind];
        //console.log(i);//
        ranks[ind] = (i-0.3)/(n+0.4);
    }
    //console.log("finish rank");//
    return ranks;
}
function makeRanked(times) {
    var len = times[0].length;
    for(var i = 0; i < times[0].length; i++) {
        if(times[1][i] == 1.0) {
            len--;
        }
    }
    var ranked = [[], [], []];
    var writes = 0;
    for(var i = 0; i < times[0].length; i++) {
        if(times[1][i] == 0.0) {
            ranked[0][writes] = times[0][i];
            ranked[1][writes] = 0;
            ranked[2][writes] = times[0].length-i;
            writes++;
            //console.log(writes);//
        }
    }
    //console.log("Finish");//
    return ranked;
}
function insertRanks(ranked, ranks) {
    ranked[1] = ranks;
}
function adjrank(ranked, n) {
    var rankadjs = [];
    for(var i = 0; i < ranked[0].length; i++) {
        var prevr;
        if(i == 0) {
            prevr = 0;
        } else {
            prevr = rankadjs[i-1];
        }
        rankadjs[i] = (ranked[2][i]*prevr+n+1)/(ranked[2][i]+1);
    }
    return rankadjs;
}
function calcXY(ranked) {
    var xy = [[], []];
    for(var c = 0; c < ranked[0].length; c++) {
        xy[0][c] = Math.log(Math.log(1/(1-ranked[1][c])));
        xy[1][c] = Math.log(ranked[0][c]);
    }
    return xy;
}
function sumXY(xy) {
    var sum = 0;
    for(var c = 0; c < xy[0].length; c++) {
        sum += xy[0][c] * xy[1][c];
    }
    return sum;
}
function sum(input) {
    var sum1 = 0;
    for(var i = 0; i < input.length; i++) {
        sum1 += input[i];
    }
    return sum1;
}
function sum2(input) {
    var sum1 = 0;
    for(var i = 0; i < input.length; i++) {
        sum1 = sum1 + input[i] * input[i];
    }
    return sum1;
}
function avg(input) {
    var sum1 = 0;
    for(var i = 0; i < input.length; i++) {
        sum1 = sum1 + input[i];
    }
    return (sum1/input.length);
}
function calcB(sumxy, sumx, sumy, sumx2, n) {
    var top = sumxy - ((sumx * sumy)/n);
    var bot = sumx2 - ((sumx * sumx)/n);
    return top/bot;
}
function calcA(avgy, b, avgx) {
    return avgy - b * avgx;
}
function calcR(sumxy, sumx, sumy, n, sumx2, sumy2) {
    var top = sumxy - ((sumx*sumy)/n);
    var bot = Math.sqrt((sumx2 - ((sumx * sumx)/n)) * (sumy2 - ((sumy * sumy)/n)));
    return top/bot;
}
function calcR2(r) {
    return r * r;
}
function calcBeta(b) {
    return 1/b;
}
function calcEta(a) {
    return Math.pow(Math.E, a);
}
function failProb(t, beta, eta) {
    var epow = -1.0*Math.pow((t/eta), beta);
    return 1.0 - Math.pow(Math.E, epow);
}
function costG(beta, times) {
    var top = 0;
    var bot = 0;
    for(i = 0; i < times[0].length; i++) {
        top += Math.pow(times[0][i], beta) * Math.log(times[0][i]);
        bot += Math.pow(times[0][i], beta);
    }
    var comfrac = top/bot;
    var sub = 0;
    for(i = 0; i < times[0].length; i++) {
        if(times[1][i] == 1) {
            continue;
        }
        sub += Math.log(times[0][i]);
    }
    sub /= 5;
    comfrac -= sub + 1/beta;
    return comfrac;
}
function gradDesc(beta, gBeta, times, num_iters) {
    num_iters = 100;
    var newBeta = beta;
    var newG = gBeta;
    for(i = 0; i < num_iters; i++) {
        newG = costG(newBeta, times);
        newBeta = newBeta - newG;
    }
    return newBeta;
}
function calcY(t, eta, beta) {
    return Math.pow(Math.E, -1*Math.pow(t / eta, beta));
}

var ETA;
var BETA;
var R2;

function input(timesIn, suspensions, t) {
    //true is suspended
    console.log("\n");
    console.log("Start");//
    var times = sort(timesIn, suspensions);
    console.log("Times");//
    var n = timesIn.length;
    console.log("N");//
    var revr = genRevr(timesIn);
    console.log("Revr");//
    var RANKED = [times[0], [], revr];
    console.log("RANKED");
    var ranked = makeRanked(times);
    console.log("Ranked");//
    var rankadjs = adjrank(ranked, n);
    console.log("Adjrank");//
    var ranks = rank(rankadjs, n);
    console.log("rank");//
    insertRanks(ranked, ranks);
    console.log("insertRank");//
    var xy = calcXY(ranked);
    console.log("xy");//
    var sumxy = sumXY(xy);
    console.log("sumxy");//
    var sumy = sum(xy[1]);
    console.log("sumy");//
    var sumy2 = sum2(xy[1]);
    console.log("sumy2");//
    var avgy = avg(xy[1]);
    console.log("avgy");//
    var sumx = sum(xy[0]);
    console.log("sumx");//
    var sumx2 = sum2(xy[0]);
    console.log("sumx2");//
    var avgx = avg(xy[0]);
    console.log("avgx");//
    var nn = ranked[0].length;
    console.log("nn");//
    var b = calcB(sumxy, sumx, sumy, sumx2, nn);
    console.log("b");//
    var a = calcA(avgy, b, avgx);
    console.log("a");//
    var r = calcR(sumxy, sumx, sumy, nn, sumx2, sumy2);
    console.log("r");//
    var r2 = calcR2(r);
    console.log("r2");//
    var beta = calcBeta(b);
    console.log("beta");//
    var eta = calcEta(a);
    console.log("eat");//
    var ft = failProb(t, beta, eta);
    console.log("ft");//
    var gb = costG(beta, times);
    console.log("gb");//
    //var gdr = gradDesc(beta, gb, times, 2, 100);
    //console.log("gdr");//

    ETA = eta;
    BETA = beta;
    R2 = r2;

    console.log("Input List: ");
    printlist(timesIn);
    printlist(suspensions);
    console.log("N: " + n);
    console.log("Sorted List: ");
    printlist(times[0]);
    printlist(times[1]);
    console.log("Unadjusted Table: ");
    printlist(RANKED[0]);
    printlist(RANKED[1]);
    console.log("Ranked Table: ");
    printlist(ranked[0]);
    printlist(ranked[1]);
    printlist(ranked[2]);
    console.log("Rank Adjustments: ");
    printlist(rankadjs);
    console.log("Ranks: ");
    printlist(ranks);
    console.log("New N: " + nn);
    console.log("XY Table: ");
    printlist(xy[0]);
    printlist(xy[1]);
    console.log("SumXY: " + sumxy);
    console.log("SumY: " + sumy);
    console.log("SumY^2: " + sumy2);
    console.log("avgY: " + avgy);
    console.log("SumX: " + sumx);
    console.log("SumX^2: " + sumx2);
    console.log("avgX: " + avgx);
    console.log("B: " + b);
    console.log("A: " + a);
    console.log("R: " + r);
    console.log("R^2: " + r2);
    console.log("Beta: " + beta);
    console.log("Eta: " + eta);
    console.log("Fail probability at time " + t + ": " + ft);
    console.log("costG of initial beta: " + gb);
    //console.log("Result of GradDesc: " + gdr);
    /**/
    return ranked;
}