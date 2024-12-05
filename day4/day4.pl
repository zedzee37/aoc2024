#!/usr/bin/perl
use strict;
use warnings;
use POSIX;

my $toMatch = "XMAS";

sub getAtPos {
    my($grid_ref, $x, $y) = @_;
    return substr(@$grid_ref[$y], $x, 1);
}

sub generateMatch {
    my ($dx, $dy) = @_;

    my $len = length $toMatch;
    $len *= 1.5;
    $len += 1 if ($len % 2 == 0); # Ensure the length is odd

    my @match = map { [( "." x $len ) =~ /./g] } (1 .. $len);
    # Initialize a 2D array of "." characters

    my $middle = int($len / 2); # Middle index (0-based)

    # Place the first character
    $match[$middle][$middle] = substr($toMatch, 0, 1);

    my $current = 1;
    my $x = $middle + $dx;
    my $y = $middle + $dy;

    # Place remaining characters
    while ($current < length($toMatch)) {
        my $char = substr($toMatch, $current, 1);
        $match[$y][$x] = $char;

        $x += $dx;
        $y += $dy;
        $current += 1;
    }

    return \@match; # Return a reference to @match
}

sub generateMatches {
    my @matches;

    for my $y (-1..1) {
       for my $x (-1..1) {
            push(@matches, generateMatch($x, $y));
       } 
    }

    return \@matches;
}

sub checkMatch {
    my ($grid_ref, $match_ref) = @_;
    my @grid = @$grid_ref;
    my @match = @$match_ref;
}

sub countMatches {
    my @grid = @_;
    my $validMatches = generateMatches;
}

open(file_handle, '<', "input.txt");

my @grid;
while (<file_handle>) {
    push(@grid, $_);
}

my $validMatches = generateMatches;