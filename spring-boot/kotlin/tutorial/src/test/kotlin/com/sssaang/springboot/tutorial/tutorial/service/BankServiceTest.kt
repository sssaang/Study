package com.sssaang.springboot.tutorial.tutorial.service

import com.sssaang.springboot.tutorial.tutorial.datasource.BankDataSource
import io.mockk.mockk
import io.mockk.verify
import org.junit.jupiter.api.Test

internal class BankServiceTest {
    private val dataSource: BankDataSource = mockk( relaxed = true )
    private val bankService = BankService(dataSource)

    @Test
    fun `should call its data source to retrieve banks`() {

        val banks = bankService.getBanks()

        verify(exactly = 1) { dataSource.retrieveBanks() }
    }
}