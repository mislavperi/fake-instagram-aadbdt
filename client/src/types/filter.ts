type Filter = {
    title?: string;
    description?: string;
    dateRange?: DateRange;
    hashtags?: string[];
    user?: string;
}

type DateRange = {
    gte: Date;
    lte: Date;
}