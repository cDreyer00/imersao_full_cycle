import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }

  getGoodbye(): string {
    return 'Goodbye World!';
  }

  getObject(): object {
    return {
      name: 'John Doe',
      age: 30
    }
  }
}
