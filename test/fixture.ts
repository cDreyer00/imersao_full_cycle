import { NestFactory } from '@nestjs/core';
import { AppModule } from '../src/app.module';
import { getDataSourceToken } from '@nestjs/typeorm';
import { DataSource } from 'typeorm';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.init();

  const dataSource = app.get<DataSource>(getDataSourceToken())
  await dataSource.synchronize(true);

  const productRepo = dataSource.getRepository('Product');

  for (let i = 0; i < 20; i++) {
    const product = productRepo.create({
      name: `Product ${i}`,
      description: `Description ${i}`,
      image_url: `https://via.placeholder.com/150?text=Product+${i}`,
      price: 1000 * Math.random(),
    });
    await productRepo.save(product);
  }

  await app.close();
}
bootstrap();