import { Request, Response, NextFunction } from "express";

const awaitHandler: (
  fn: any
) => (req: Request, res: Response, next: NextFunction) => Promise<void> = (
  fn: any
) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    try {
      await fn(req, res, next);
    } catch (err) {
      next(err);
    }
  };
};

export default awaitHandler;
